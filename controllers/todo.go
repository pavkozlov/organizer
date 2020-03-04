package controllers

import (
	"github.com/pavkozlov/organizer/config"
	"github.com/pavkozlov/organizer/models"
)

func GetAllTodos(todo *[]models.Todo) (err error) {
	if err = config.Db.Find(todo).Error; err != nil {
		return err
	}
	return nil
}
