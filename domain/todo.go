package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/router"
)

func Todo(
	app *fiber.App,
	todoController controller.TodoController,
	todolistController controller.TodolistController,
	attachmentController controller.AttachmentController,
) {
	router.TodoRouter(
		app,
		todoController,
		todolistController,
		attachmentController,
	)
}
