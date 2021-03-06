package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/middleware"
)

func AdminRouter(app *fiber.App, controller controller.AdminController) {
	app.Get("/admin/users", middleware.AuthenticatedWithRole("admin"), controller.GetAllRegisteredUser)
}
