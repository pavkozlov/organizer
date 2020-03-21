package todo

import (
	"github.com/pavkozlov/organizer/organizer"
)

func ListTodo(todo *[]Todo) (err error) {
	if err = organizer.Db.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func RetrieveTodo(todo *Todo, id string) (err error) {
	if err = organizer.Db.Where("id = ?", id).First(&todo).Error; err != nil {
		return err
	}
	return nil
}

func DestroyTodo(todo *Todo, id string) (err error) {
	if err = organizer.Db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {
	if err = organizer.Db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *Todo, new_title, id string) (err error) {
	if err = organizer.Db.Model(&todo).Where("id = ?", id).Update("title", new_title).Find(&todo).Error; err != nil {
		return err
	}
	return nil
}
