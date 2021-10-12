package repository

import (
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type TodoRepository interface {
	FindAll() ([]model.Todo, error)
	FindById(id uuid.UUID) (model.Todo, error)
	Create(todo model.Todo) error
	UpdateById(id uuid.UUID, todo model.Todo) error
	DeleteById(id uuid.UUID) error
}
