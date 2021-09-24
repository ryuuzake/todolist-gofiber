package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/service"
)

type AdminControllerImpl struct {
	Service service.AdminService
}

func (controller AdminControllerImpl) GetAllRegisteredUser(ctx *fiber.Ctx) error {
	// TODO: Check Role

	users, err := controller.Service.GetAllRegisteredUser()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(users)
}
