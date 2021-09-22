package service

import (
	"errors"
	"strings"

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

func (service *AuthServiceImpl) LoginUser(email string, password string) (model.User, error) {
	user, err := service.Repository.FindByEmail(email)
	if err != nil {
		return model.User{}, errors.New("user not found")
	}

	// TODO: Check User password hash
	if strings.Compare(user.Password, password) != 0 {
		return model.User{}, errors.New("user creds failed")
	}

	return user, nil
}
