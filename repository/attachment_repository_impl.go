package repository

import (
	"errors"
	"github.com/ryuuzake/todolist-gofiber/model"
)

type AttachmentRepositoryInMemoryImpl struct {
	Attachments []model.Attachment
}

func (repo *AttachmentRepositoryInMemoryImpl) FindById(id int) (model.Attachment, error) {
	for _, attachment := range repo.Attachments {
		if attachment.Id == id {
			return attachment, nil
		}
	}

	return model.Attachment{}, errors.New("attachment not found")
}

func (repo *AttachmentRepositoryInMemoryImpl) FindAllFromTodolistId(id int) ([]model.Attachment, error) {
	attachments := make([]model.Attachment, 0)
	for _, attachment := range repo.Attachments {
		if attachment.TodolistId == id {
			attachments = append(attachments, attachment)
		}
	}

	return attachments, nil
}

func (repo *AttachmentRepositoryInMemoryImpl) Create(attachment model.Attachment) error {
	repo.Attachments = append(repo.Attachments, attachment)

	return nil
}

func (repo *AttachmentRepositoryInMemoryImpl) UpdateById(id int, newAttachment model.Attachment) error {
	for i, attachment := range repo.Attachments {
		if attachment.Id == id {
			p := &repo.Attachments[i]
			*p = newAttachment
			return nil
		}
	}

	return errors.New("todolist not found")
}

func (repo *AttachmentRepositoryInMemoryImpl) DeleteById(id int) error {
	for i, attachment := range repo.Attachments {
		if attachment.Id == id {
			repo.Attachments = append(repo.Attachments[:i], repo.Attachments[i+1:]...)
			return nil
		}
	}

	return errors.New("attachment not found")
}
