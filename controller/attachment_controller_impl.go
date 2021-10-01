package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/service"
)

type AttachmentControllerImpl struct {
	Service service.AttachmentService
}

func (controller AttachmentControllerImpl) GetAllAttachment(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (controller AttachmentControllerImpl) GetByIdAttachment(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (controller AttachmentControllerImpl) CreateAttachment(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (controller AttachmentControllerImpl) UpdateAttachment(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (controller AttachmentControllerImpl) DeleteAttachment(ctx *fiber.Ctx) error {
	panic("implement me")
}
