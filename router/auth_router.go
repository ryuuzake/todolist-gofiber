package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
)

func AuthRouter(app *fiber.App, controller controller.AuthController) {
	app.Post("/auth/register", controller.RegisterUser)
	app.Post("/auth/login", controller.LoginUser)
	app.Post("/auth/forget-password", controller.ForgetPassword)
	app.Post("/auth/reset-password/:token", controller.ResetPassword)
	app.Post("/auth/logout", controller.LogOut)
}
