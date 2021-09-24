package service

import "github.com/ryuuzake/todolist-gofiber/model"

type AdminService interface {
	GetAllRegisteredUser() ([]model.User, error)
}
