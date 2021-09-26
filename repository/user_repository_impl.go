package repository

import (
	"errors"
	"strings"
	"time"

	"github.com/ryuuzake/todolist-gofiber/model"
)

type UserRepositoryInMemoryImpl struct {
	Users []model.User
}

func (repo *UserRepositoryInMemoryImpl) FindAll() ([]model.User, error) {
	return repo.Users, nil
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
	// Add Default Option
	user.Id = cap(repo.Users) + 1
	user.RoleId = 2 // Normal User Role Id
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	repo.Users = append(repo.Users, user)

	return nil
}
