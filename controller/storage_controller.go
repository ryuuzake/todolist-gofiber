package controller

import "github.com/gofiber/fiber/v2"

type StorageController interface {
	GetFile(ctx *fiber.Ctx) error
}
