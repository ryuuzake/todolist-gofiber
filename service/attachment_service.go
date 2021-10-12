package service

import (
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type AttachmentService interface {
	GetAllAttachmentWithTodolistId(id uuid.UUID) ([]model.Attachment, error)
	GetAttachmentById(id uuid.UUID) (model.Attachment, error)
	CreateAttachmentWithTodolistId(id uuid.UUID, attachment model.Attachment) error
	UpdateAttachmentById(id uuid.UUID, attachment model.Attachment) error
	DeleteAttachmentById(id uuid.UUID) error
}
