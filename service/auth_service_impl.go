package service

import (
	"errors"

	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type AuthServiceImpl struct {
	Repository repository.UserRepository
}

func (service *AuthServiceImpl) RegisterUser(email, password string) error {
	if _, err := service.Repository.FindByEmail(email); err == nil {
		return errors.New("user already registered")
	}

	err := service.Repository.Create(model.User{Email: email, Password: password})
	if err != nil {
		return err
	}

	return nil
}
