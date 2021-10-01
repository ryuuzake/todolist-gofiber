package service

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type TodoServiceImpl struct {
	Repository repository.TodoRepository
}

func (service TodoServiceImpl) GetAllTodo() ([]model.Todo, error) {
	todos, err := service.Repository.FindAll()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (service TodoServiceImpl) GetTodoById(id int) (model.Todo, error) {
	todo, err := service.Repository.FindById(id)

	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func (service TodoServiceImpl) CreateTodo(todo model.Todo) error {
	if err := service.Repository.Create(todo); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) UpdateTodo(id int, todo model.Todo) error {
	if err := service.Repository.UpdateById(id, todo); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) DeleteTodo(id int) error {
	if err := service.Repository.DeleteById(id); err != nil {
		return err
	}

	return nil
}
