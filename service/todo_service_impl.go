package service

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	Repository     repository.TodolistRepository
}

func (service TodoServiceImpl) GetAll() ([]model.Todo, error) {
	todos, err := service.TodoRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (service TodoServiceImpl) GetById(id int) (model.Todo, error) {
	todo, err := service.TodoRepository.FindById(id)

	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func (service TodoServiceImpl) Create(todo model.Todo) error {
	if err := service.TodoRepository.Create(todo); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) Update(id int, todo model.Todo) error {
	if err := service.TodoRepository.UpdateById(id, todo); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) Delete(id int) error {
	if err := service.TodoRepository.DeleteById(id); err != nil {
		return err
	}

	return nil
}
