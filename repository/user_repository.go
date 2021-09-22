package repository

import "github.com/ryuuzake/todolist-gofiber/model"

type UserRepository interface {
	FindByEmail(email string) (model.User, error)
	Create(model.User) error
}
