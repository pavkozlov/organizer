package account

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/organizer"
	"net/http"
	"time"
)

func Register(ctx *gin.Context) {
	username, password := ctx.PostForm("username"), ctx.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid username/password"})
		return
	}

	salt := generateRandomString(64)
	user := User{
		Username: username,
		Salt:     salt,
		Password: encryptPassword(password, salt),
	}

	if err := saveUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func Login(ctx *gin.Context) {
	user := User{}
	username, password := ctx.PostForm("username"), ctx.PostForm("password")
	if getUserByUsername(&user, username) != nil || !authorize(username, password) {
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

func RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.PostForm("refreshToken")
	if len(refreshToken) != 128 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad token"})
		return
	}

	refreshTokenSql := refreshTokenRaw{}
	if err := getRefreshToken(&refreshTokenSql, refreshToken); err != nil {
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
