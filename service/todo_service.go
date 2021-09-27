package service

import "github.com/ryuuzake/todolist-gofiber/model"

type TodoService interface {
	GetAll() ([]model.Todo, error)
	GetById(id int) (model.Todo, error)
	Create(todo model.Todo) error
	Update(id int, todo model.Todo) error
	Delete(id int) error
}
