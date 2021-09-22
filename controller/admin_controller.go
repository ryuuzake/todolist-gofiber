package controller

import "github.com/gofiber/fiber/v2"

type AdminController interface {
	GetAllRegisteredUser(ctx *fiber.Ctx) error
}
