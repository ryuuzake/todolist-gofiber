package controller

import "github.com/gofiber/fiber/v2"

type TodoController interface {
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error

	GetAllTodolist(ctx *fiber.Ctx) error
	GetByIdTodolist(ctx *fiber.Ctx) error
	CreateTodolist(ctx *fiber.Ctx) error
	UpdateTodolist(ctx *fiber.Ctx) error
	DeleteTodolist(ctx *fiber.Ctx) error

	GetAllPhoto(ctx *fiber.Ctx) error
	GetByIdPhoto(ctx *fiber.Ctx) error
	CreatePhoto(ctx *fiber.Ctx) error
	UpdatePhoto(ctx *fiber.Ctx) error
	DeletePhoto(ctx *fiber.Ctx) error
}
