package service

import (
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type TodoService interface {
	GetAllTodo() ([]model.Todo, error)
	GetTodoById(id uuid.UUID) (model.Todo, error)
	CreateTodo(todo model.Todo) error
	UpdateTodo(id uuid.UUID, todo model.Todo) error
	DeleteTodo(id uuid.UUID) error
}
