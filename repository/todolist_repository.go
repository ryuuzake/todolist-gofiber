package repository

import (
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type TodolistRepository interface {
	FindAll() ([]model.Todolist, error)
	FindAllWithTodoId(id uuid.UUID) ([]model.Todolist, error)
	FindById(id uuid.UUID) (model.Todolist, error)
	Create(model.Todolist) error
	UpdateById(id uuid.UUID, todolist model.Todolist) error
	DeleteById(id uuid.UUID) error
}
