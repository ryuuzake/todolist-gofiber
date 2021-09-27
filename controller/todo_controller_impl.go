package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/service"
	"github.com/ryuuzake/todolist-gofiber/validator"
)

type TodoControllerImpl struct {
	Service service.TodoService
}

func (controller TodoControllerImpl) GetAll(ctx *fiber.Ctx) error {
	todos, err := controller.Service.GetAll()

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

	todo, err := controller.Service.GetById(todoId)

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

	// TODO: Error Handling with GoFiber
	// TODO: Add UserId to TodoPayload
	if ok := controller.Service.Create(model.Todo{Title: todoPayload.Title}); ok != nil {
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
	if err := controller.Service.Update(todoId, todo); err != nil {
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

	if err := controller.Service.Delete(todolistId); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (controller TodoControllerImpl) UploadPhoto(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller TodoControllerImpl) UpdatePhoto(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
