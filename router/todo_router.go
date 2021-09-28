package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/middleware"
)

func TodoRouter(app *fiber.App, controller controller.TodoController) {
	app.Get("/todos", middleware.Authenticated(), controller.GetAll)
	app.Get("/todos/:todoId", middleware.Authenticated(), controller.GetById)
	app.Post("/todos", middleware.Authenticated(), controller.Create)
	app.Patch("/todos/:todoId", middleware.Authenticated(), controller.Update)
	app.Delete("/todos/:todoId", middleware.Authenticated(), controller.Delete)
	app.Post("/todos/:todoId/todolists", middleware.Authenticated(), controller.CreateTodolist)
	app.Patch("/todos/:todoId/todolists/:todolistId", middleware.Authenticated(), controller.UpdateTodolist)
	app.Post("/todos/:todoId/todolists/:todolistId/photos", middleware.Authenticated(), controller.UploadPhoto)
	app.Patch("/todos/:todoId/todolists/:todolistId/photos/:photoId", middleware.Authenticated(), controller.UpdatePhoto)
}
