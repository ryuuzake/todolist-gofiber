package repository

import "github.com/ryuuzake/todolist-gofiber/model"

type TodoRepository interface {
	FindAll() ([]model.Todo, error)
	FindById(id int) (model.Todo, error)
	Create(todo model.Todo) error
	UpdateById(id int, todo model.Todo) error
	DeleteById(id int) error
}
