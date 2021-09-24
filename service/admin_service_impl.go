package service

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type AdminServiceImpl struct {
	Repository repository.UserRepository
}

func (service AdminServiceImpl) GetAllRegisteredUser() ([]model.User, error) {
	users, err := service.Repository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
