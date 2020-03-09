package middleware

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/config"
	"github.com/pavkozlov/organizer/models"
	"net/http"
	"strings"
)

func SetCors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Next()
}

func CustomBasicAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !authenticateUser(pair[0], pair[1]) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func authenticateUser(username, password string) bool {
	var user models.User
	err := config.Db.Where(models.User{Username: username, Password: password}).First(&user)
	if err.Error != nil {
		return false
	}
	return true

}
