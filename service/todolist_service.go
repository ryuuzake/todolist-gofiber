package service

import "github.com/ryuuzake/todolist-gofiber/model"

type TodolistService interface {
	GetAll() ([]model.Todolist, error)
}
