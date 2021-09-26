package repository

import (
	"errors"

	"github.com/ryuuzake/todolist-gofiber/model"
)

type RoleRepositoryInMemoryImpl struct {
	Roles []model.Role
}

func (repo *RoleRepositoryInMemoryImpl) FindById(id int) (model.Role, error) {
	for _, role := range repo.Roles {
		if role.Id == id {
			return role, nil
		}
	}

	return model.Role{}, errors.New("role not found")
}
