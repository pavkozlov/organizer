package views

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/models"
	"github.com/pavkozlov/organizer/settings"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Register(ctx *gin.Context) {
	// Проверка входных параметров
	// ToDo вынести это в middleware
	username, password := ctx.PostForm("username"), ctx.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Укажите username + password"})
		return
	}

	// Создание соли
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	salt := ""
	for i := 0; i < 64; i++ {
		salt += string(charset[seededRand.Intn(len(charset))])
	}

	// хеширование пароля
	sha_512 := sha512.New()
	sha_512.Write([]byte(password + salt))
	encryptedPassword := sha_512.Sum([]byte(""))

	// формирование юзера
	user := models.User{
		Username: username,
		Salt:     salt,
		Password: hex.EncodeToString(encryptedPassword),
	}

	// проверка на успешность создания
	// ToDo вынести это в контроллеры
	if settings.Db.Save(&user).Error == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь " + username + " уже существует"})
	}

}

func Login(ctx *gin.Context) {

	username, password := ctx.PostForm("username"), ctx.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u := models.User{}
	db := settings.Db
	e := db.Where("username = ?", username).Find(&u)

	fmt.Println(u, e == nil)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": ctx.PostForm("username"),
		"id":       u.ID,
		"expired":  time.Now().Local().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(settings.SecretKey))

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
