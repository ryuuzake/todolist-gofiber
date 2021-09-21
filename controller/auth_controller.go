package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	RegisterUser(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
	ForgetPassword(ctx *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error
	LogOut(ctx *fiber.Ctx) error
}
