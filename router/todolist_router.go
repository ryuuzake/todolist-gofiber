package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
)

func TodolistRouter(app *fiber.App, controller controller.TodolistController) {
	app.Get("/todolists", controller.GetAll)
	app.Get("/todolists/:todolistId", controller.GetById)
	app.Post("/todolists", controller.Create)
	app.Patch("/todolists/:todolistId", controller.Update)
	app.Delete("/todolists/:todolistId", controller.Delete)
	app.Post("/todolists/:todolistId/photos", controller.UploadPhoto)
	app.Patch("/todolists/:todolistId/photos/:photoId", controller.UpdatePhoto)
}
