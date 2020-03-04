package views

import (
	"github.com/gin-gonic/gin"
	"github.com/pavkozlov/organizer/config"
	"github.com/pavkozlov/organizer/controllers"
	"github.com/pavkozlov/organizer/models"
	"net/http"
)

// GetTodos func return all items
func GetTodos(ctx *gin.Context) {
	todo := make([]models.Todo, 0)
	err := controllers.GetAllTodos(&todo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": todo})
	}
}

// GetATodo func return item by id
func GetATodo(ctx *gin.Context) {

	todo := models.Todo{}
	config.Db.Where("id = ?", ctx.Params.ByName("id")).First(&todo)
	ctx.JSON(http.StatusOK, todo)
}

// DeleteTodo func delete item by id
func DeleteTodo(ctx *gin.Context) {
	todo := make([]models.Todo, 0)
	config.Db.Where("id = ?", ctx.Params.ByName("id")).Delete(&todo)
	ctx.JSON(http.StatusNoContent, gin.H{"t": todo})
}

// CreateATodo func create item
func CreateATodo(ctx *gin.Context) {

	var title string
	if title = ctx.PostForm("title"); len(title) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Укажите title"})
		return
	}

	todo := models.Todo{Title: title}
	config.Db.Create(&todo)

	ctx.JSON(http.StatusCreated, gin.H{"data": todo})
}

// Update an existing Todo
func UpdateATodo(ctx *gin.Context) {
	var title string
	if title = ctx.PostForm("title"); len(title) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Укажите title"})
		return
	}
	todo := models.Todo{}
	config.Db.Model(&todo).Where("id = ?", ctx.Params.ByName("id")).Update("title", title).Find(&todo)
	ctx.JSON(http.StatusOK, todo)
}
