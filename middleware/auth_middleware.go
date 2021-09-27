package middleware

import (
	"github.com/ryuuzake/todolist-gofiber/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticated() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   config.Config.JWTSecret,
		ErrorHandler: jwtError,
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func AuthenticatedWithRole(role string) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: config.Config.JWTSecret,
		// TODO: Weird Handling of passing parameter role
		SuccessHandler: func(c *fiber.Ctx) error {
			return jwtCheckRole(c, role)
		},
		ErrorHandler: jwtError,
	})
}

func jwtCheckRole(ctx *fiber.Ctx, role string) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	roleClaim := claims["role"]

	if roleClaim == nil || strings.Compare(roleClaim.(string), role) != 0 {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	return ctx.Next()
}
