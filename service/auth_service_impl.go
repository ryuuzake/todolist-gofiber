package service

import (
	"errors"
	"github.com/ryuuzake/todolist-gofiber/helper"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
	"strings"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func (service *AuthServiceImpl) RegisterUser(email, password string) error {
	if _, err := service.UserRepository.FindByEmail(email); err == nil {
		return errors.New("user already registered")
	}

	err := service.UserRepository.Create(model.User{Email: email, Password: password})
	if err != nil {
		return err
	}

	return nil
}

func (service *AuthServiceImpl) LoginUser(email string, password string) (string, error) {
	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	role, err := service.RoleRepository.FindById(user.RoleId)
	if err != nil {
		return "", errors.New("role not found")
	}

	user.Role = role

	// TODO: Check User password hash
	if strings.Compare(user.Password, password) != 0 {
		return "", errors.New("user creds failed")
	}

	return helper.EncodeJWT(user)
}
