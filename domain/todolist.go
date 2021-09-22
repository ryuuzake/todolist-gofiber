package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/router"
)

func Todolist(app *fiber.App, controller controller.TodolistController) {
	router.TodolistRouter(app, controller)
}
