package repository

import (
	"errors"

	"github.com/ryuuzake/todolist-gofiber/model"
)

type TodolistRepositoryInMemoryImpl struct {
	Todolists []model.Todolist
}

func (repo *TodolistRepositoryInMemoryImpl) FindAll() ([]model.Todolist, error) {
	return repo.Todolists, nil
}

func (repo *TodolistRepositoryInMemoryImpl) FindById(id int) (model.Todolist, error) {
	for _, todolist := range repo.Todolists {
		if todolist.Id == id {
			return todolist, nil
		}
	}

	return model.Todolist{}, errors.New("todolist not found")
}

func (repo *TodolistRepositoryInMemoryImpl) Create(todolist model.Todolist) error {
	repo.Todolists = append(repo.Todolists, todolist)

	return nil
}

func (repo *TodolistRepositoryInMemoryImpl) UpdateById(id int, newTodolist model.Todolist) error {
	for i, todolist := range repo.Todolists {
		if todolist.Id == id {
			p := &repo.Todolists[i]
			*p = newTodolist
			return nil
		}
	}

	return errors.New("todolist not found")
}

func (repo *TodolistRepositoryInMemoryImpl) DeleteById(id int) error {
	for i, todolist := range repo.Todolists {
		if todolist.Id == id {
			repo.Todolists = append(repo.Todolists[:i], repo.Todolists[i+1:]...)
			return nil
		}
	}

	return errors.New("todolist not found")
}
