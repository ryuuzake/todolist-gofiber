package service

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
	"time"
)

type TodoServiceImpl struct {
	TodoRepository       repository.TodoRepository
	TodolistRepository   repository.TodolistRepository
	AttachmentRepository repository.AttachmentRepository
}

func (service TodoServiceImpl) GetAllTodo() ([]model.Todo, error) {
	todos, err := service.TodoRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (service TodoServiceImpl) GetTodoById(id int) (model.Todo, error) {
	todo, err := service.TodoRepository.FindById(id)

	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func (service TodoServiceImpl) CreateTodo(todo model.Todo) error {
	if err := service.TodoRepository.Create(todo); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) UpdateTodo(id int, todo model.Todo) error {
	if err := service.TodoRepository.UpdateById(id, todo); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) DeleteTodo(id int) error {
	if err := service.TodoRepository.DeleteById(id); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) CreateTodolistWithTodoId(id int, todolist model.Todolist) error {
	todolist.TodoId = id

	if err := service.TodolistRepository.Create(todolist); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) UpdateTodolistWithTodoId(id int, todolist model.Todolist) error {
	if err := service.TodolistRepository.UpdateById(id, todolist); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) CreateAttachmentWithTodolistId(id int, attachment model.Attachment) error {
	// TODO: Default value place in service or repository
	attachment.TodolistId = id
	attachment.CreatedAt = time.Now()
	attachment.UpdatedAt = time.Now()

	if err := service.AttachmentRepository.Create(attachment); err != nil {
		return err
	}

	return nil
}

func (service TodoServiceImpl) UpdateAttachmentWithId(id int, attachment model.Attachment) error {
	attachment.UpdatedAt = time.Now()

	if err := service.AttachmentRepository.UpdateById(id, attachment); err != nil {
		return err
	}

	return nil
}
