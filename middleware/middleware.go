package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetCors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Next()
}
