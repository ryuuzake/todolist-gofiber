package controller

import (
	"github.com/ryuuzake/todolist-gofiber/model"
	"time"

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

	user := model.User{
		FullName: registerUserPayload.FullName,
		Email:    registerUserPayload.Email,
		Password: registerUserPayload.Password,
	}
	ok := controller.Service.RegisterUser(user)

	// TODO: Error Handling with GoFiber
	if ok != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": ok.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller *AuthControllerImpl) LoginUser(ctx *fiber.Ctx) error {
	loginUserPayload := new(validator.LoginUserPayload)

	// TODO: Error Handling with GoFiber
	if err := ctx.BodyParser(loginUserPayload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validator.ValidateLoginUserPayload(loginUserPayload)
	// TODO: Error Handling with GoFiber
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErr)
	}

	token, err := controller.Service.LoginUser(
		loginUserPayload.Email,
		loginUserPayload.Password,
	)

	// TODO: Error Handling with GoFiber
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"access_token": token,
		"token_type":   "Bearer",
		// TODO: Remove Duplication, probably by creating Token struct type
		"expired_in": time.Now().Add(time.Hour * 72).Unix(),
	})
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
