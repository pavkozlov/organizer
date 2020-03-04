package urls

import (
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/middleware"
	"github.com/pavkozlov/organizer/views"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetCors)
	// r.Use(cors.Default())
	v1 := r.Group("/v1")
	{
		v1.GET("todo", views.GetTodos)
		v1.GET("todo/:id", views.GetATodo)
		v1.DELETE("todo/:id", views.DeleteTodo)
		v1.POST("todo", views.CreateATodo)
		v1.PUT("todo/:id", views.UpdateATodo)
	}
	return r
}
