package service

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type TodolistServiceImpl struct {
	Repository repository.TodolistRepository
}

func (service TodolistServiceImpl) GetAllTodolistWithTodoId(id int) ([]model.Todolist, error) {
	todolists, err := service.Repository.FindAllWithTodoId(id)

	if err != nil {
		return nil, err
	}

	return todolists, nil
}

func (service TodolistServiceImpl) GetTodolistById(id int) (model.Todolist, error) {
	todolist, err := service.Repository.FindById(id)

	if err != nil {
		return model.Todolist{}, err
	}

	return todolist, nil
}

func (service TodolistServiceImpl) CreateTodolistWithTodoId(id int, todolist model.Todolist) error {
	todolist.TodoId = id

	if err := service.Repository.Create(todolist); err != nil {
		return err
	}

	return nil
}

func (service TodolistServiceImpl) UpdateTodolistById(id int, todolist model.Todolist) error {
	if err := service.Repository.UpdateById(id, todolist); err != nil {
		return err
	}

	return nil
}

func (service TodolistServiceImpl) DeleteTodolistById(id int) error {
	if err := service.Repository.DeleteById(id); err != nil {
		return err
	}

	return nil
}
