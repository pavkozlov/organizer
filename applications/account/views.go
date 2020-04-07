package account

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/organizer"
)

// Эндпоинт для регистрации
func Register(ctx *gin.Context) {
	json := userForm{}
	if e := ctx.ShouldBind(&json); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid username/password"})
		return
	}

	salt := generateRandomString(64)
	user := User{
		Username: json.Username,
		Salt:     salt,
		Password: encryptPassword(json.Password, salt),
	}

	if err := createUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

// Эндпоинт для входа
func Login(ctx *gin.Context) {
	user := User{}
	json := userForm{}

	if e := ctx.ShouldBind(&json); e != nil || !authorize(json.Username, json.Password, &user) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid username/password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"expired":  time.Now().Add(time.Minute * 30).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(organizer.SecretKey))

	userAgent := ctx.GetHeader("User-Agent")
	session := Sessions{UserID: user.ID, UserAgent: userAgent}

	if err := getOrCrateRefreshToken(&session, generateRandomString(128)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accessToken": tokenString, "refreshToken": session.RefreshToken})
}

// Эндпоинт для обновления токена
func RefreshToken(ctx *gin.Context) {

	json := refreshToken{}
	if e := ctx.ShouldBind(&json); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad token"})
		return
	}

	refreshTokenSql := refreshTokenRaw{}
	if err := getRefreshToken(&refreshTokenSql, json.RefreshToken); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if refreshTokenSql.ExpiresIn.Unix() <= time.Now().Unix() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Token expired"})
		return
	}

	if err := deleteRefreshToken(refreshTokenSql.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRt := Sessions{RefreshToken: generateRandomString(128), UserID: refreshTokenSql.UserID, UserAgent: ctx.GetHeader("User-Agent")}
	if err := createRefreshToken(&newRt); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  generateToken(refreshTokenSql.Username, refreshTokenSql.UserID),
		"refreshToken": newRt.RefreshToken,
	})

}
