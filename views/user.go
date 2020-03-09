package views

import (
	"crypto/sha512"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/config"
	"github.com/pavkozlov/organizer/models"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// return items list
func CreateUsers(ctx *gin.Context) {

	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	salt := make([]byte, 64)
	for i := range salt {
		salt[i] = charset[seededRand.Intn(len(charset))]
	}

	sha_512 := sha512.New()
	sha_512.Write([]byte(ctx.PostForm("password") + string(salt)))

	user := models.User{
		Username: ctx.PostForm("username"),
		Salt:     string(salt),
		Password: base64.URLEncoding.EncodeToString(sha_512.Sum([]byte(""))),
	}

	config.Db.Save(&user)

	ctx.JSON(http.StatusOK, user)
}
