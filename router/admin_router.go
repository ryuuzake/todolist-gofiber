package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
)

func AdminRouter(app *fiber.App, controller controller.AdminController) {
	app.Get("/admin/users", controller.GetAllRegisteredUser)
}
