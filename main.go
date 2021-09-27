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
		Users: []model.User{
			{Id: 1, Email: "admin@example.com", Password: "password", RoleId: 1},
		},
	}

	roleRepository := &repository.RoleRepositoryInMemoryImpl{
		Roles: []model.Role{
			{Id: 1, Name: "admin", Description: "Admin Role for User"},
			{Id: 2, Name: "normal", Description: "Normal User"},
		},
	}

	domain.Admin(app, &controller.AdminControllerImpl{
		Service: &service.AdminServiceImpl{
			Repository: userRepository,
		},
	})

	domain.Todo(app, &controller.TodoControllerImpl{
		Service: &service.TodoServiceImpl{
			Repository: &repository.TodolistRepositoryInMemoryImpl{
				Todolists: make([]model.Todolist, 0),
			},
			TodoRepository: &repository.TodoRepositoryInMemoryImpl{
				Todos: make([]model.Todo, 0),
			},
		},
	})

	domain.Auth(app, &controller.AuthControllerImpl{
		Service: &service.AuthServiceImpl{
			UserRepository: userRepository,
			RoleRepository: roleRepository,
		},
	})

	app.Listen("localhost:3000")
}
