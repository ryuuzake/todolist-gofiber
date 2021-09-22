package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/domain"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
	"github.com/ryuuzake/todolist-gofiber/service"
)

func main() {
	app := fiber.New()

	domain.Auth(app, &controller.AuthControllerImpl{
		Service: &service.AuthServiceImpl{
			Repository: &repository.UserRepositoryInMemoryImpl{
				Users: make([]model.User, 0),
			},
		},
	})

	app.Listen("localhost:3000")
}
