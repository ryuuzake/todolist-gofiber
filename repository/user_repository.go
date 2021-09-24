package repository

import "github.com/ryuuzake/todolist-gofiber/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	Create(model.User) error
}
