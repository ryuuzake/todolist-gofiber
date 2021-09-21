package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/domain"
)

func main() {
	app := fiber.New()

	domain.Auth(app, &controller.AuthControllerImpl{})

	app.Listen(":3000")
}
