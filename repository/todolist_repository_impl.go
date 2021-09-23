package repository

import "github.com/ryuuzake/todolist-gofiber/model"

type TodolistRepositoryInMemoryImpl struct {
	Todolists []model.Todolist
}

func (repository TodolistRepositoryInMemoryImpl) FindAll() ([]model.Todolist, error) {
	return repository.Todolists, nil
}
