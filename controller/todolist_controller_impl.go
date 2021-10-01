package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/service"
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
	panic("implement me")
}

func (controller *TodolistControllerImpl) UpdateTodolist(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (controller *TodolistControllerImpl) DeleteTodolist(ctx *fiber.Ctx) error {
	panic("implement me")
}
