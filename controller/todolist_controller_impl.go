package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/service"
	"github.com/ryuuzake/todolist-gofiber/validator"
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
	todolistId, err := strconv.Atoi(ctx.Params("todolistId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	todolist, err := controller.Service.GetById(todolistId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(todolist)
}

func (controller TodolistControllerImpl) Create(ctx *fiber.Ctx) error {
	createTodolistPayload := new(validator.CreateTodolistPayload)

	// TODO: Error Handling with GoFiber
	if err := ctx.BodyParser(createTodolistPayload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validator.ValidateCreateTodolistPayload(createTodolistPayload)
	// TODO: Error Handling with GoFiber
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErr)
	}

	// TODO: Error Handling with GoFiber
	if ok := controller.Service.Create(model.Todolist{Task: createTodolistPayload.Task}); ok != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": ok,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller TodolistControllerImpl) Update(ctx *fiber.Ctx) error {
	todolistId, err := strconv.Atoi(ctx.Params("todolistId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	updateTodolistPayload := new(validator.UpdateTodolistPayload)

	// TODO: Error Handling with GoFiber
	if err := ctx.BodyParser(updateTodolistPayload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// TODO: Refactor this
	todolistModel := model.Todolist{Task: updateTodolistPayload.Task}
	if err := controller.Service.Update(todolistId, todolistModel); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller TodolistControllerImpl) Delete(ctx *fiber.Ctx) error {
	todolistId, err := strconv.Atoi(ctx.Params("todolistId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := controller.Service.Delete(todolistId); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (controller TodolistControllerImpl) UploadPhoto(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller TodolistControllerImpl) UpdatePhoto(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
