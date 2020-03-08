package views

import (
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/controllers"
	"github.com/pavkozlov/organizer/models"
	"net/http"
)

// return items list
func GetTodosList(ctx *gin.Context) {
	todo := make([]models.Todo, 0)

	if err := controllers.GetAllTodos(&todo); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": todo})
	}
}

// return 1 item
func GetTodo(ctx *gin.Context) {
	todo := models.Todo{}
	id := ctx.Params.ByName("id")

	if err := controllers.GetTodo(&todo, id); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": todo})
	}
}

// return 204
func DeleteTodo(ctx *gin.Context) {
	todo := models.Todo{}
	id := ctx.Params.ByName("id")

	if err := controllers.DeleteTodo(&todo, id); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{"data": todo})
	}

}

// return created todo
func CreateATodo(ctx *gin.Context) {
	title := ctx.PostForm("title")
	todo := models.Todo{Title: title}

	if len(title) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := controllers.CreateTodo(&todo); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"data": todo})
	}
}

// return updated todo
func UpdateATodo(ctx *gin.Context) {
	todo := models.Todo{}
	id := ctx.Params.ByName("id")
	title := ctx.PostForm("title")

	if len(title) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := controllers.UpdateTodo(&todo, title, id); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": todo})
	}
}
