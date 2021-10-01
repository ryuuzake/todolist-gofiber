package service

import "github.com/ryuuzake/todolist-gofiber/model"

type AttachmentService interface {
	GetAllAttachmentWithTodolistId(id int) ([]model.Attachment, error)
	GetAttachmentById(id int) (model.Attachment, error)
	CreateAttachmentWithTodolistId(id int, attachment model.Attachment) error
	UpdateAttachmentById(id int, attachment model.Attachment) error
	DeleteAttachmentById(id int) error
}
