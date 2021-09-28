package repository

import "github.com/ryuuzake/todolist-gofiber/model"

type AttachmentRepository interface {
	FindById(id int) (model.Attachment, error)
	FindAllFromTodolistId(id int) ([]model.Attachment, error)
	Create(attachment model.Attachment) error
	UpdateById(id int, attachment model.Attachment) error
	DeleteById(id int) error
}
