package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ryuuzake/todolist-gofiber/config"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type UserRepositoryPostgresImpl struct {
	*sqlx.DB
}

func NewUserRepositoryPostgresImpl() *UserRepositoryPostgresImpl {
	db, err := config.OpenDBConnection()
	if err != nil {
		panic(err.Error())
	}

	return &UserRepositoryPostgresImpl{DB: db}
}

func (repo *UserRepositoryPostgresImpl) FindAll() ([]model.User, error) {
	var users []model.User

	query := `SELECT * FROM users`

	err := repo.Get(&users, query)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (repo *UserRepositoryPostgresImpl) FindByEmail(email string) (model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE email=$1 LIMIT 1`

	err := repo.Get(&user, query, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryPostgresImpl) Create(user model.User) error {
	query := `INSERT INTO users (full_name, email, password, role_id)
		VALUES(:full_name, :email, :password, :role_id)`

	_, err := repo.NamedExec(query, user)
	if err != nil {
		return err
	}

	return nil
}
