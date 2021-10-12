package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ryuuzake/todolist-gofiber/config"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type TodoRepositoryPostgresImpl struct {
	*sqlx.DB
}

func NewTodoRepositoryPostgresImpl() *TodoRepositoryPostgresImpl {
	db, err := config.OpenDBConnection()
	if err != nil {
		panic(err.Error())
	}

	return &TodoRepositoryPostgresImpl{DB: db}
}

func (repo *TodoRepositoryPostgresImpl) FindAll() ([]model.Todo, error) {
	todos := make([]model.Todo, 0)

	query := `SELECT * FROM todos`

	err := repo.Select(&todos, query)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (repo *TodoRepositoryPostgresImpl) FindById(id uuid.UUID) (model.Todo, error) {
	var todo model.Todo

	query := `SELECT * FROM todos WHERE id=$1 LIMIT 1`

	err := repo.Get(&todo, query, id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (repo *TodoRepositoryPostgresImpl) Create(todo model.Todo) error {
	query := `INSERT INTO todos (user_id, title)
		VALUES(:user_id, :title)`

	_, err := repo.NamedExec(query, todo)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TodoRepositoryPostgresImpl) UpdateById(id uuid.UUID, todo model.Todo) error {
	// Set Id to Id from params
	todo.Id = id

	query := `UPDATE todos SET title=:title WHERE id=:id`

	_, err := repo.NamedExec(query, todo)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TodoRepositoryPostgresImpl) DeleteById(id uuid.UUID) error {
	query := `DELETE FROM todos WHERE id=$1`

	_, err := repo.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
