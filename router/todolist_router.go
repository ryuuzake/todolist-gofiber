package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/middleware"
)

func TodolistRouter(app *fiber.App, controller controller.TodolistController) {
	app.Get("/todolists", middleware.Authenticated(), controller.GetAll)
	app.Get("/todolists/:todolistId", middleware.Authenticated(), controller.GetById)
	app.Post("/todolists", middleware.Authenticated(), controller.Create)
	app.Patch("/todolists/:todolistId", middleware.Authenticated(), controller.Update)
	app.Delete("/todolists/:todolistId", middleware.Authenticated(), controller.Delete)
	app.Post("/todolists/:todolistId/photos", middleware.Authenticated(), controller.UploadPhoto)
	app.Patch("/todolists/:todolistId/photos/:photoId", middleware.Authenticated(), controller.UpdatePhoto)
}
