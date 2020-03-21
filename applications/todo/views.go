package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodoList(ctx *gin.Context) {
	todo := []Todo{}
	if err := ListTodo(&todo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": todo})

}

func GetTodo(ctx *gin.Context) {
	todo := Todo{}
	id := ctx.Params.ByName("id")

	if err := RetrieveTodo(&todo, id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": todo})

}

func DeleteTodo(ctx *gin.Context) {
	todo := Todo{}
	id := ctx.Params.ByName("id")

	if err := DestroyTodo(&todo, id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"data": todo})

}

func CreateATodo(ctx *gin.Context) {
	title := ctx.PostForm("title")
	todo := Todo{Title: title}

	if len(title) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Title len must be > 0"})
		return
	}

	if err := CreateTodo(&todo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": todo})

}

func UpdateATodo(ctx *gin.Context) {
	todo := Todo{}
	id := ctx.Params.ByName("id")
	title := ctx.PostForm("title")

	if len(title) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Title len must be > 0"})
		return
	}

	if err := UpdateTodo(&todo, title, id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": todo})

}
