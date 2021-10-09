package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ryuuzake/todolist-gofiber/config"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type RoleRepositoryPostgresImpl struct {
	*sqlx.DB
}

func NewRoleRepositoryPostgresImpl() *RoleRepositoryPostgresImpl {
	db, err := config.OpenDBConnection()
	if err != nil {
		panic(err.Error())
	}

	return &RoleRepositoryPostgresImpl{DB: db}
}

func (repo *RoleRepositoryPostgresImpl) FindById(id uuid.UUID) (model.Role, error) {
	var role model.Role

	query := `SELECT * FROM roles WHERE id=$1 LIMIT 1`

	err := repo.Get(&role, query, id)
	if err != nil {
		return role, err
	}

	return role, nil
}
