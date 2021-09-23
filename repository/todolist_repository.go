package repository

import "github.com/ryuuzake/todolist-gofiber/model"

type TodolistRepository interface {
	FindAll() ([]model.Todolist, error)
}
