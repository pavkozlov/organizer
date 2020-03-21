package views

import (
	"crypto/sha512"
	"encoding/hex"
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
	// Создание соли
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	salt := ""
	for i := 0; i < 64; i++ {
		salt += string(charset[seededRand.Intn(len(charset))])
	}

	// хеширование пароля
	sha_512 := sha512.New()
	sha_512.Write([]byte(ctx.PostForm("password") + salt))
	encryptedPassword := sha_512.Sum([]byte(""))

	// формирование юзера
	user := models.User{
		Username: ctx.PostForm("username"),
		Salt:     salt,
		Password: hex.EncodeToString(encryptedPassword),
	}

	// ToDo вынести это в контроллеры
	if settings.Db.Save(&user).Error == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь " + ctx.PostForm("password") + " уже существует"})
	}

}

func authorize(username, password string) bool {
	user := models.User{}
	e := settings.Db.Where("username = ?", username).Find(&user)
	if e.Error != nil {
		return false
	}

	sha_512 := sha512.New()
	sha_512.Write([]byte(password + user.Salt))
	encryptedPassword := sha_512.Sum([]byte(""))
	stringPassword := hex.EncodeToString(encryptedPassword)

	if user.Password == stringPassword {
		return true
	} else {
		return false
	}

}

func Login(ctx *gin.Context) {
	u := models.User{}
	db := settings.Db
	db.Where("username = ?", ctx.PostForm("username")).Find(&u)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": ctx.PostForm("username"),
		"id":       u.ID,
		"expired":  time.Now().Local().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(settings.SecretKey))

	auth := authorize(ctx.PostForm("username"), ctx.PostForm("password"))

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString, "correct": auth})
}
