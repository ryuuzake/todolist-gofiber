package repository

import "github.com/ryuuzake/todolist-gofiber/model"

type TodolistRepository interface {
	FindAll() ([]model.Todolist, error)
	FindById(id int) (model.Todolist, error)
	Create(model.Todolist) error
	UpdateById(id int, todolist model.Todolist) error
	DeleteById(id int) error
}
