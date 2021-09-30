package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/middleware"
)

func TodoRouter(
	app *fiber.App,
	todoController controller.TodoController,
	todolistController controller.TodolistController,
	attachmentController controller.AttachmentController,
) {
	todoRoute := app.Group("/todos", middleware.Authenticated())
	todoRoute.Get("", todoController.GetAll)
	todoRoute.Get("/:todoId", todoController.GetById)
	todoRoute.Post("", todoController.Create)
	todoRoute.Patch("/:todoId", todoController.Update)
	todoRoute.Delete("/:todoId", todoController.Delete)

	todolistRoute := todoRoute.Group("/:todoId/todolists")
	todolistRoute.Get("", todolistController.GetAllTodolist)
	todolistRoute.Get("/:todolistId", todolistController.GetByIdTodolist)
	todolistRoute.Post("", todolistController.CreateTodolist)
	todolistRoute.Patch("/:todolistId", todolistController.UpdateTodolist)
	todolistRoute.Delete("/:todolistId", todolistController.DeleteTodolist)

	photoRoute := todolistRoute.Group("/:todolistId/photos")
	photoRoute.Get("", attachmentController.GetAllAttachment)
	photoRoute.Get("/:photoId", attachmentController.GetByIdAttachment)
	photoRoute.Post("", attachmentController.CreateAttachment)
	photoRoute.Patch("/:photoId", attachmentController.UpdateAttachment)
	photoRoute.Delete("/:photoId", attachmentController.DeleteAttachment)
}
