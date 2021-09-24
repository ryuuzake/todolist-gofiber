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

	userRepository := &repository.UserRepositoryInMemoryImpl{
		Users: make([]model.User, 0),
	}

	domain.Admin(app, &controller.AdminControllerImpl{
		Service: &service.AdminServiceImpl{
			Repository: userRepository,
		},
	})

	domain.Todolist(app, &controller.TodolistControllerImpl{
		Service: &service.TodolistServiceImpl{
			Repository: &repository.TodolistRepositoryInMemoryImpl{
				Todolists: make([]model.Todolist, 0),
			},
		},
	})

	domain.Auth(app, &controller.AuthControllerImpl{
		Service: &service.AuthServiceImpl{
			Repository: userRepository,
		},
	})

	app.Listen("localhost:3000")
}
