package repository

import (
	"errors"
	"github.com/ryuuzake/todolist-gofiber/model"
	"time"
)

type TodoRepositoryInMemoryImpl struct {
	Todos []model.Todo
}

func (repo *TodoRepositoryInMemoryImpl) FindAll() ([]model.Todo, error) {
	return repo.Todos, nil
}

func (repo *TodoRepositoryInMemoryImpl) FindById(id int) (model.Todo, error) {
	for _, todo := range repo.Todos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return model.Todo{}, errors.New("todo not found")
}

func (repo *TodoRepositoryInMemoryImpl) Create(todo model.Todo) error {
	// Add Default Option
	todo.Id = cap(repo.Todos) + 1
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	repo.Todos = append(repo.Todos, todo)

	return nil
}

func (repo *TodoRepositoryInMemoryImpl) UpdateById(id int, newTodo model.Todo) error {
	for i, todo := range repo.Todos {
		if todo.Id == id {
			// TODO: Simplify or try to abstract it
			p := &repo.Todos[i]
			(*p).Title = newTodo.Title
			(*p).UpdatedAt = time.Now()
			return nil
		}
	}

	return errors.New("todo not found")
}

func (repo *TodoRepositoryInMemoryImpl) DeleteById(id int) error {
	for i, todo := range repo.Todos {
		if todo.Id == id {
			repo.Todos = append(repo.Todos[:i], repo.Todos[i+1:]...)
			return nil
		}
	}

	return errors.New("todo not found")
}
