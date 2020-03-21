package account

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/settings"
	"net/http"
	"strings"
	"time"
)

func Register(ctx *gin.Context) {
	salt := generateSalt(64)

	username, password := ctx.PostForm("username"), ctx.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := User{
		Username: username,
		Salt:     salt,
		Password: password,
	}

	if err := saveUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, user)
	}

}

func Login(ctx *gin.Context) {
	user := User{}
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if getUserByUsername(&user, username) != nil || authorize(username, password) == false {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid username/password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"expired":  time.Now().Local().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(settings.SecretKey))

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Auth(ctx *gin.Context) {

	var tokenString string
	auth := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)
	if len(auth) == 2 && auth[0] == "Bearer" {
		tokenString = auth[1]
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// todo add validate expired < now
		ctx.JSON(http.StatusOK, gin.H{
			"id":       claims["id"],
			"username": claims["username"],
			"expired":  claims["expired"],
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad token"})
	}

}
