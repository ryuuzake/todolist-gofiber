package service

import "github.com/ryuuzake/todolist-gofiber/model"

type AuthService interface {
	RegisterUser(user model.User) error
	LoginUser(email, password string) (string, error)
}
