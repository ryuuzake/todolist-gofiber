package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/router"
)

func Todo(app *fiber.App, controller controller.TodoController) {
	router.TodoRouter(app, controller)
}
