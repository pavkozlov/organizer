package urls

import (
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/middleware"
	"github.com/pavkozlov/organizer/views"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetCors)

	//r.Use(middleware.CustomBasicAuth())

	todo := r.Group("/todo")
	{
		todo.GET("/", views.GetTodosList)
		todo.GET("/:id", views.GetTodo)
		todo.DELETE("/:id", views.DeleteTodo)
		todo.POST("/", views.CreateATodo)
		todo.PUT("/:id", views.UpdateATodo)
	}
	user := r.Group("/user")
	{
		user.POST("/", views.CreateUsers)
	}
	return r
}
