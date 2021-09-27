package service

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type TodoServiceImpl struct {
	Repository repository.TodolistRepository
}

func (service TodoServiceImpl) GetAll() ([]model.Todolist, error) {
	todolists, err := service.Repository.FindAll()

	if err != nil {
		return nil, err
	}

	return todolists, nil
}

func (service TodoServiceImpl) GetById(id int) (model.Todolist, error) {
	todolist, err := service.Repository.FindById(id)

	if err != nil {
		return model.Todolist{}, err
	}

	return todolist, nil
}

func (service TodoServiceImpl) Create(todolist model.Todolist) error {
	if err := service.Repository.Create(todolist); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) Update(id int, todolist model.Todolist) error {
	if err := service.Repository.UpdateById(id, todolist); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) Delete(id int) error {
	if err := service.Repository.DeleteById(id); err != nil {
		return err
	}

	return nil
}
