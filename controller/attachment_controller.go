package controller

import "github.com/gofiber/fiber/v2"

type AttachmentController interface {
	GetAllAttachment(ctx *fiber.Ctx) error
	GetByIdAttachment(ctx *fiber.Ctx) error
	CreateAttachment(ctx *fiber.Ctx) error
	UpdateAttachment(ctx *fiber.Ctx) error
	DeleteAttachment(ctx *fiber.Ctx) error
}
