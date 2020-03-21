package account

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/organizer"
	"net/http"
	"strings"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Bearer" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := auth[1]
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(organizer.SecretKey), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid || int64(claims["expired"].(float64))-time.Now().Unix() <= 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad token"})
		}

		user := gin.H{
			"id":       claims["id"],
			"username": claims["username"],
		}

		ctx.Set("user", user)
		ctx.Next()

	}
}
