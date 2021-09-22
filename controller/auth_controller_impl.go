package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/service"
	"github.com/ryuuzake/todolist-gofiber/validator"
)

type AuthControllerImpl struct {
	Service service.AuthService
}

func (controller *AuthControllerImpl) RegisterUser(ctx *fiber.Ctx) error {
	registerUserPayload := new(validator.RegisterUserPayload)

	// TODO: Error Handling with GoFiber
	// https://docs.gofiber.io/guide/error-handling
	if err := ctx.BodyParser(registerUserPayload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validator.ValidateRegisterUserPayload(registerUserPayload)
	// TODO: Error Handling with GoFiber
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErr)
	}

	ok := controller.Service.RegisterUser(
		registerUserPayload.Email,
		registerUserPayload.Password,
	)

	// TODO: Error Handling with GoFiber
	if ok != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": ok,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller *AuthControllerImpl) LoginUser(ctx *fiber.Ctx) error {
	return nil
}

func (controller *AuthControllerImpl) ForgetPassword(ctx *fiber.Ctx) error {
	return nil
}

func (controller *AuthControllerImpl) ResetPassword(ctx *fiber.Ctx) error {
	return nil
}

func (controller *AuthControllerImpl) LogOut(ctx *fiber.Ctx) error {
	return nil
}
