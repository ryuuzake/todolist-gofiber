package service

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
	"time"
)

type AttachmentServiceImpl struct {
	Repository repository.AttachmentRepository
}

func (service AttachmentServiceImpl) GetAllAttachmentWithTodolistId(id int) ([]model.Attachment, error) {
	panic("implement me")
}

func (service AttachmentServiceImpl) GetAttachmentById(id int) (model.Attachment, error) {
	panic("implement me")
}

func (service AttachmentServiceImpl) CreateAttachmentWithTodolistId(id int, attachment model.Attachment) error {
	// TODO: Default value place in service or repository
	attachment.TodolistId = id
	attachment.CreatedAt = time.Now()
	attachment.UpdatedAt = time.Now()

	if err := service.Repository.Create(attachment); err != nil {
		return err
	}

	return nil
}

func (service AttachmentServiceImpl) UpdateAttachmentById(id int, attachment model.Attachment) error {
	attachment.UpdatedAt = time.Now()

	if err := service.Repository.UpdateById(id, attachment); err != nil {
		return err
	}

	return nil
}

func (service AttachmentServiceImpl) DeleteAttachmentById(id int) error {
	panic("implement me")
}
