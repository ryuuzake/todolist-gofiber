package repository

import "github.com/ryuuzake/todolist-gofiber/model"

type RoleRepository interface {
	FindById(id int) (model.Role, error)
}
