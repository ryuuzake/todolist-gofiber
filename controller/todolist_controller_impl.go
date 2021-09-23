package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/service"
)

type TodolistControllerImpl struct {
	Service service.TodolistService
}

func (controller TodolistControllerImpl) GetAll(ctx *fiber.Ctx) error {
	todolists, err := controller.Service.GetAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(todolists)
}

func (controller TodolistControllerImpl) GetById(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller TodolistControllerImpl) Create(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller TodolistControllerImpl) Update(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller TodolistControllerImpl) Delete(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller TodolistControllerImpl) UploadPhoto(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller TodolistControllerImpl) UpdatePhoto(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
