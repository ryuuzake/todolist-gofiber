package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/middleware"
)

func TodoRouter(app *fiber.App, controller controller.TodoController) {
	todoRoute := app.Group("/todos", middleware.Authenticated())
	todoRoute.Get("", controller.GetAll)
	todoRoute.Get("/:todoId", controller.GetById)
	todoRoute.Post("", controller.Create)
	todoRoute.Patch("/:todoId", controller.Update)
	todoRoute.Delete("/:todoId", controller.Delete)

	todolistRoute := todoRoute.Group("/:todoId/todolists")
	todolistRoute.Get("", controller.GetAllTodolist)
	todolistRoute.Get("/:todolistId", controller.GetByIdTodolist)
	todolistRoute.Post("", controller.CreateTodolist)
	todolistRoute.Patch("/:todolistId", controller.UpdateTodolist)
	todolistRoute.Delete("/:todolistId", controller.DeleteTodolist)

	photoRoute := todolistRoute.Group("/:todolistId/photos")
	photoRoute.Get("", controller.GetAllPhoto)
	photoRoute.Get("/:photoId", controller.GetByIdPhoto)
	photoRoute.Post("", controller.CreatePhoto)
	photoRoute.Patch("/:photoId", controller.UpdatePhoto)
	photoRoute.Delete("/:photoId", controller.DeletePhoto)
}
