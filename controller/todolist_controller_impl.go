package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/service"
	"github.com/ryuuzake/todolist-gofiber/validator"
)

type TodolistControllerImpl struct {
	Service service.TodolistService
}

func (controller *TodolistControllerImpl) GetAllTodolist(ctx *fiber.Ctx) error {
	todoId, err := uuid.FromString(ctx.Params("todoId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	todolists, err := controller.Service.GetAllTodolistWithTodoId(todoId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(todolists)
}

func (controller *TodolistControllerImpl) GetByIdTodolist(ctx *fiber.Ctx) error {
	todolistId, err := uuid.FromString(ctx.Params("todolistId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	todolist, err := controller.Service.GetTodolistById(todolistId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(todolist)
}

func (controller *TodolistControllerImpl) CreateTodolist(ctx *fiber.Ctx) error {
	todoId, err := uuid.FromString(ctx.Params("todoId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	todolistPayload := new(validator.TodolistPayload)
	if err := ctx.BodyParser(todolistPayload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validator.ValidateCreateTodolistPayload(todolistPayload)
	if validationErr != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(validationErr)
	}

	todolist := model.Todolist{TodoId: todoId, StatusId: todolistPayload.StatusId, Task: todolistPayload.Task}
	if ok := controller.Service.CreateTodolistWithTodoId(todoId, todolist); ok != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": ok,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

func (controller *TodolistControllerImpl) UpdateTodolist(ctx *fiber.Ctx) error {
	todoId, err := uuid.FromString(ctx.Params("todoId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	todolistId, err := uuid.FromString(ctx.Params("todolistId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	todolistPayload := new(validator.TodolistPayload)
	if err := ctx.BodyParser(todolistPayload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validator.ValidateUpdateTodolistPayload(todolistPayload)
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErr)
	}

	todolist := model.Todolist{TodoId: todoId, StatusId: todolistPayload.StatusId, Task: todolistPayload.Task}
	if err := controller.Service.UpdateTodolistById(todolistId, todolist); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller *TodolistControllerImpl) DeleteTodolist(ctx *fiber.Ctx) error {
	todolistId, err := uuid.FromString(ctx.Params("todolistId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	if err := controller.Service.DeleteTodolistById(todolistId); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
