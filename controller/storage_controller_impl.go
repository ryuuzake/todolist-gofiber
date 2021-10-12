package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/service"
)

type StorageControllerImpl struct {
	Service service.FileUploadService
}

func (controller *StorageControllerImpl) GetFile(ctx *fiber.Ctx) error {
	fileName := ctx.Params("filename")

	filePath, err := controller.Service.RetrieveFile(fileName)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "file not found",
		})
	}

	return ctx.SendFile(filePath)
}
