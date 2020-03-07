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

func GetTodo(todo *models.Todo, id string) (err error) {
	if err = config.Db.Where("id = ?", id).First(&todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo *models.Todo, id string) (err error) {
	if err = config.Db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *models.Todo) (err error) {
	if err = config.Db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *models.Todo, new_title, id string) (err error) {
	if err = config.Db.Model(&todo).Where("id = ?", id).Update("title", new_title).Find(&todo).Error; err != nil {
		return err
	}
	return nil
}
