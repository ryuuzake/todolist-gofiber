package service

import "github.com/ryuuzake/todolist-gofiber/model"

type TodolistService interface {
	GetAllTodolistWithTodoId(id int) ([]model.Todolist, error)
	GetTodolistById(id int) (model.Todolist, error)
	CreateTodolistWithTodoId(id int, todolist model.Todolist) error
	UpdateTodolistById(id int, todolist model.Todolist) error
	DeleteTodolistById(id int) error
}
