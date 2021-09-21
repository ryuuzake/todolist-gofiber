package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/router"
)

func Auth(app *fiber.App, authController controller.AuthController) {
	router.AuthRouter(app, authController)
}
