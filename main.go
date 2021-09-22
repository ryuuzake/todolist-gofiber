package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/domain"
	"github.com/ryuuzake/todolist-gofiber/service"
)

func main() {
	app := fiber.New()

	domain.Auth(app, &controller.AuthControllerImpl{
		Service: &service.AuthServiceImpl{},
	})

	app.Listen(":3000")
}
