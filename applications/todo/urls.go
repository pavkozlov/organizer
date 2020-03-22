package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/applications/account"
)

func SetupRouter(r *gin.Engine) {
	todoRouter := r.Group("/todo")
	todoRouter.Use(account.JWTAuth())
	{
		todoRouter.GET("/", GetTodoList)
		todoRouter.GET("/:id", GetTodo)
		todoRouter.DELETE("/:id", DeleteTodo)
		todoRouter.POST("/", CreateATodo)
		todoRouter.PUT("/:id", UpdateATodo)
	}
}
