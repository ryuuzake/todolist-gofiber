package repository

import (
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type AttachmentRepository interface {
	FindById(id uuid.UUID) (model.Attachment, error)
	FindAllFromTodolistId(id uuid.UUID) ([]model.Attachment, error)
	Create(attachment model.Attachment) error
	UpdateById(id uuid.UUID, attachment model.Attachment) error
	DeleteById(id uuid.UUID) error
}
