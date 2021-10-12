package service

import (
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
	"time"
)

type AttachmentServiceImpl struct {
	Repository repository.AttachmentRepository
}

func (service AttachmentServiceImpl) GetAllAttachmentWithTodolistId(id uuid.UUID) ([]model.Attachment, error) {
	attachments, err := service.Repository.FindAllFromTodolistId(id)

	if err != nil {
		return nil, err
	}

	return attachments, nil
}

func (service AttachmentServiceImpl) GetAttachmentById(id uuid.UUID) (model.Attachment, error) {
	attachment, err := service.Repository.FindById(id)

	if err != nil {
		return model.Attachment{}, err
	}

	return attachment, nil
}

func (service AttachmentServiceImpl) CreateAttachmentWithTodolistId(id uuid.UUID, attachment model.Attachment) error {
	// TODO: Default value place in service or repository
	attachment.TodolistId = id
	attachment.CreatedAt = time.Now()
	attachment.UpdatedAt = time.Now()

	if err := service.Repository.Create(attachment); err != nil {
		return err
	}

	return nil
}

func (service AttachmentServiceImpl) UpdateAttachmentById(id uuid.UUID, attachment model.Attachment) error {
	attachment.UpdatedAt = time.Now()

	if err := service.Repository.UpdateById(id, attachment); err != nil {
		return err
	}

	return nil
}

func (service AttachmentServiceImpl) DeleteAttachmentById(id uuid.UUID) error {
	if err := service.Repository.DeleteById(id); err != nil {
		return err
	}

	return nil
}
