package service

import (
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type TodolistService interface {
	GetAllTodolistWithTodoId(id uuid.UUID) ([]model.Todolist, error)
	GetTodolistById(id uuid.UUID) (model.Todolist, error)
	CreateTodolistWithTodoId(id uuid.UUID, todolist model.Todolist) error
	UpdateTodolistById(id uuid.UUID, todolist model.Todolist) error
	DeleteTodolistById(id uuid.UUID) error
}
