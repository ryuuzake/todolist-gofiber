package service

import "github.com/ryuuzake/todolist-gofiber/model"

type TodoService interface {
	GetAll() ([]model.Todolist, error)
	GetById(id int) (model.Todolist, error)
	Create(todolist model.Todolist) error
	Update(id int, todolist model.Todolist) error
	Delete(id int) error
}
