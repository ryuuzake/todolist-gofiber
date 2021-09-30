package controller

import (
	"strconv"

	"github.com/ryuuzake/todolist-gofiber/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/service"
	"github.com/ryuuzake/todolist-gofiber/validator"
)

type TodoControllerImpl struct {
	Service service.TodoService
}

func (controller TodoControllerImpl) GetAll(ctx *fiber.Ctx) error {
	todos, err := controller.Service.GetAllTodo()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(todos)
}

func (controller TodoControllerImpl) GetById(ctx *fiber.Ctx) error {
	todoId, err := strconv.Atoi(ctx.Params("todoId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	todo, err := controller.Service.GetTodoById(todoId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(todo)
}

func (controller TodoControllerImpl) Create(ctx *fiber.Ctx) error {
	todoPayload := new(validator.TodoPayload)

	// TODO: Error Handling with GoFiber
	if err := ctx.BodyParser(todoPayload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validator.ValidateCreateTodoPayload(todoPayload)
	// TODO: Error Handling with GoFiber
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErr)
	}

	// TODO: This assume user can only create todo for themself, ask if need to be otherwise
	userClaim := helper.DecodeJWT(ctx)
	todo := model.Todo{UserId: userClaim.UserId, Title: todoPayload.Title}
	// TODO: Error Handling with GoFiber
	if ok := controller.Service.CreateTodo(todo); ok != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": ok,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller TodoControllerImpl) Update(ctx *fiber.Ctx) error {
	todoId, err := strconv.Atoi(ctx.Params("todoId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	todoPayload := new(validator.TodoPayload)

	// TODO: Error Handling with GoFiber
	if err := ctx.BodyParser(todoPayload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validator.ValidateUpdateTodoPayload(todoPayload)
	// TODO: Error Handling with GoFiber
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErr)
	}

	// TODO: Refactor this
	todo := model.Todo{Title: todoPayload.Title}
	if err := controller.Service.UpdateTodo(todoId, todo); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller TodoControllerImpl) Delete(ctx *fiber.Ctx) error {
	todolistId, err := strconv.Atoi(ctx.Params("todoId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := controller.Service.DeleteTodo(todolistId); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
