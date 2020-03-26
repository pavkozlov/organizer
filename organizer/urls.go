package organizer

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(SetCors())
	r.Use(cors.Default())
	return r
}
