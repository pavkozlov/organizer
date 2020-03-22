package account

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/settings"
	"net/http"
	"time"
)

func Register(ctx *gin.Context) {
	salt := generateRandomString(64)

	username, password := ctx.PostForm("username"), ctx.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := User{
		Username: username,
		Salt:     salt,
		Password: encryptPassword(password, salt),
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
	userAgent := ctx.GetHeader("User-Agent")

	if getUserByUsername(&user, username) != nil || !authorize(username, password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid username/password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"expired":  time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(settings.SecretKey))

	s := Sessions{}
	settings.Db.Where(Sessions{UserID: user.ID, UserAgent: userAgent}).Attrs(Sessions{RefreshToken: generateRandomString(128)}).FirstOrCreate(&s)

	ctx.JSON(http.StatusOK, gin.H{"accessToken": tokenString, "refreshToken": s.RefreshToken})
}
