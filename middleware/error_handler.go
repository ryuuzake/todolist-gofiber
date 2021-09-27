package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	// Retrieve the custom status code if it's a fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message
	}

	// Send custom error JSON Structure
	err = ctx.Status(code).JSON(fiber.Map{"message": msg})

	// Return from handler
	return nil
}
