package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/config"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/domain"
	"github.com/ryuuzake/todolist-gofiber/middleware"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
	"github.com/ryuuzake/todolist-gofiber/service"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	userRepository := repository.NewUserRepositoryPostgresImpl()
	roleRepository := repository.NewRoleRepositoryPostgresImpl()
	todoRepository := repository.NewTodoRepositoryPostgresImpl()

	domain.Admin(app, &controller.AdminControllerImpl{
		Service: &service.AdminServiceImpl{
			Repository: userRepository,
		},
	})

	domain.Todo(
		app,
		&controller.TodoControllerImpl{
			Service: &service.TodoServiceImpl{
				Repository: todoRepository,
			},
		},
		&controller.TodolistControllerImpl{
			Service: &service.TodolistServiceImpl{
				Repository: &repository.TodolistRepositoryInMemoryImpl{
					Todolists: make([]model.Todolist, 0),
				},
			},
		},
		&controller.AttachmentControllerImpl{
			AttachmentService: &service.AttachmentServiceImpl{
				Repository: &repository.AttachmentRepositoryInMemoryImpl{
					Attachments: make([]model.Attachment, 0),
				},
			},
			FileUploadService: &service.FileUploadServiceImpl{},
		},
	)

	domain.Auth(app, &controller.AuthControllerImpl{
		Service: &service.AuthServiceImpl{
			UserRepository: userRepository,
			RoleRepository: roleRepository,
		},
	})

	log.Fatal(app.Listen(config.Config.Address))
}
