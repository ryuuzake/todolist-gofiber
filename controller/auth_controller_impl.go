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

	errors := validator.ValidateRegisterUserPayload(registerUserPayload)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	controller.Service.RegisterUser(
		registerUserPayload.Email,
		registerUserPayload.Password,
	)

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
