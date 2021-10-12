package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/config"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/domain"
	"github.com/ryuuzake/todolist-gofiber/middleware"
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
	todolistRepository := repository.NewTodolistRepositoryPostgresImpl()
	attachmentRepository := repository.NewAttachmentRepositoryPostgresImpl()

	fileUploadService := service.FileUploadServiceImpl{}

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
				Repository: todolistRepository,
			},
		},
		&controller.AttachmentControllerImpl{
			AttachmentService: &service.AttachmentServiceImpl{
				Repository: attachmentRepository,
			},
			FileUploadService: fileUploadService,
		},
	)

	domain.Auth(app, &controller.AuthControllerImpl{
		Service: &service.AuthServiceImpl{
			UserRepository: userRepository,
			RoleRepository: roleRepository,
		},
	})

	domain.Storage(app, &controller.StorageControllerImpl{
		Service: fileUploadService,
	})

	log.Fatal(app.Listen(config.Config.Address))
}
