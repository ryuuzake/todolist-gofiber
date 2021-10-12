package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
	"github.com/ryuuzake/todolist-gofiber/router"
)

func Storage(
	app *fiber.App,
	storageController controller.StorageController,
) {
	router.StorageRouter(
		app,
		storageController,
	)
}
