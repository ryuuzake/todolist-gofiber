package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/service"
	"github.com/ryuuzake/todolist-gofiber/validator"
	"strconv"
)

type TodolistControllerImpl struct {
	Service service.TodolistService
}

func (controller *TodolistControllerImpl) GetAllTodolist(ctx *fiber.Ctx) error {
	todoId, err := strconv.Atoi(ctx.Params("todoId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	todos, err := controller.Service.GetAllTodolistWithTodoId(todoId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(todos)
}

func (controller *TodolistControllerImpl) GetByIdTodolist(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (controller *TodolistControllerImpl) CreateTodolist(ctx *fiber.Ctx) error {
	todoId, err := strconv.Atoi(ctx.Params("todoId"))
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

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller *TodolistControllerImpl) UpdateTodolist(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (controller *TodolistControllerImpl) DeleteTodolist(ctx *fiber.Ctx) error {
	panic("implement me")
}
