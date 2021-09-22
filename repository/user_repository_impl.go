package repository

import (
	"errors"
	"strings"

	"github.com/ryuuzake/todolist-gofiber/model"
)

type UserRepositoryInMemoryImpl struct {
	Users []model.User
}

func (repo *UserRepositoryInMemoryImpl) FindByEmail(email string) (model.User, error) {
	for _, user := range repo.Users {
		if strings.Compare(user.Email, email) == 0 {
			return user, nil
		}
	}

	return model.User{}, errors.New("user not found")
}

func (repo *UserRepositoryInMemoryImpl) Create(user model.User) error {
	repo.Users = append(repo.Users, user)

	return nil
}
