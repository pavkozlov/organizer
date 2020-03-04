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
	todo := r.Group("/todo")
	{
		todo.GET("/", views.GetTodos)
		todo.GET("/:id", views.GetATodo)
		todo.DELETE("/:id", views.DeleteTodo)
		todo.POST("/", views.CreateATodo)
		todo.PUT("/:id", views.UpdateATodo)
	}
	return r
}
