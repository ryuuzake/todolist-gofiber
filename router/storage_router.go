package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/controller"
)

func StorageRouter(
	app *fiber.App,
	storageController controller.StorageController,
) {
	app.Get("/storage/:filename", storageController.GetFile)
}
