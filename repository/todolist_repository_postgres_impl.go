package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ryuuzake/todolist-gofiber/config"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type TodolistRepositoryPostgresImpl struct {
	*sqlx.DB
}

func NewTodolistRepositoryPostgresImpl() *TodolistRepositoryPostgresImpl {
	db, err := config.OpenDBConnection()
	if err != nil {
		panic(err.Error())
	}

	return &TodolistRepositoryPostgresImpl{DB: db}
}

func (repo *TodolistRepositoryPostgresImpl) FindAll() ([]model.Todolist, error) {
	todolists := make([]model.Todolist, 0)

	query := `SELECT * FROM todolists`

	err := repo.Select(&todolists, query)
	if err != nil {
		return todolists, err
	}

	return todolists, nil
}

func (repo *TodolistRepositoryPostgresImpl) FindAllWithTodoId(id uuid.UUID) ([]model.Todolist, error) {
	todolists := make([]model.Todolist, 0)

	query := `SELECT * FROM todolists WHERE todo_id=$1`

	err := repo.Select(&todolists, query, id)
	if err != nil {
		return todolists, err
	}

	return todolists, nil
}

func (repo *TodolistRepositoryPostgresImpl) FindById(id uuid.UUID) (model.Todolist, error) {
	var todolist model.Todolist

	query := `SELECT * FROM todolists WHERE id=$1 LIMIT 1`

	err := repo.Get(&todolist, query, id)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (repo *TodolistRepositoryPostgresImpl) Create(todolist model.Todolist) error {
	query := `INSERT INTO todolists (todo_id, task, status_id)
		VALUES(:todo_id, :task, :status_id)`

	_, err := repo.NamedExec(query, todolist)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TodolistRepositoryPostgresImpl) UpdateById(id uuid.UUID, todolist model.Todolist) error {
	// Set Id to Id from params
	todolist.Id = id

	query := `UPDATE todolists SET todo_id=:todo_id, status_id=:status_id, task=:task WHERE id=:id`

	_, err := repo.NamedExec(query, todolist)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TodolistRepositoryPostgresImpl) DeleteById(id uuid.UUID) error {
	query := `DELETE FROM todolists WHERE id=$1`

	_, err := repo.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
