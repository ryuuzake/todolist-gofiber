package repository

import (
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type RoleRepository interface {
	FindById(id uuid.UUID) (model.Role, error)
}
