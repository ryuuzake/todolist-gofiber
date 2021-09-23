package service

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type TodolistServiceImpl struct {
	Repository repository.TodolistRepository
}

func (service TodolistServiceImpl) GetAll() ([]model.Todolist, error) {
	todolists, err := service.Repository.FindAll()

	if err != nil {
		return nil, err
	}

	return todolists, nil
}
