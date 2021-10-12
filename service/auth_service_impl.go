package service

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/helper"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func (service *AuthServiceImpl) RegisterUser(user model.User) error {
	if _, err := service.UserRepository.FindByEmail(user.Email); err == nil {
		return errors.New("user already registered")
	}

	hashPass, err := helper.GenerateHashFromPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPass

	// TODO: Get Normal Role
	user.RoleId, _ = uuid.FromString("7c41a513-7606-42af-81f6-0fa0d0a1676e")

	err = service.UserRepository.Create(user)
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

	err = helper.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return "", errors.New("user creds failed")
	}

	return helper.EncodeJWT(user)
}
