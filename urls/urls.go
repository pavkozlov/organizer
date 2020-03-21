package urls

import (
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/applications/account"
	"github.com/pavkozlov/organizer/middleware"
	"github.com/pavkozlov/organizer/views"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetCors)

	todo := r.Group("/todo")
	todo.Use(middleware.JWTAuth())
	{
		todo.GET("/", views.GetTodosList)
		todo.GET("/:id", views.GetTodo)
		todo.DELETE("/:id", views.DeleteTodo)
		todo.POST("/", views.CreateATodo)
		todo.PUT("/:id", views.UpdateATodo)
	}

	user := r.Group("/user")
	{
		user.POST("/login", account.Login)
		user.POST("/reg", account.Register)
		user.POST("/refresh", account.RefreshToken)
	}
	return r
}
