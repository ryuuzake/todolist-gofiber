package service

import "github.com/ryuuzake/todolist-gofiber/model"

type AuthService interface {
	RegisterUser(email, password string) error
	LoginUser(email, password string) (model.User, error)
}
