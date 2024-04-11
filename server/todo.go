package server

import (
	"simple-todo/db"
	"simple-todo/model"
)

var d = db.DbInit()

func GetList() []model.Todo {
	var todo []model.Todo
	d.Where("completed = ?", false).Find(&todo)
	return todo
}
func GetAllList() []model.Todo {
	var todo []model.Todo
	d.Find(&todo)
	return todo
}

func AddTodo(content string) uint {
	todo := model.Todo{
		Content: content,
	}
	d.Create(&todo)
	return uint(todo.ID)
}
func EditTodo(id uint) uint {
	var change bool
	d.Model(&model.Todo{}).Where("id = ?", id).Select("completed").Scan(&change)
	d.Model(&model.Todo{}).Where("id = ?", id).Update("completed", !change)
	return id
}
func DeleteTodo(id uint) bool {
	d.Delete(&model.Todo{}, id)
	return true
}
