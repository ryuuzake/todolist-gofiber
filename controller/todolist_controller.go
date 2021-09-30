package controller

import "github.com/gofiber/fiber/v2"

type TodolistController interface {
	GetAllTodolist(ctx *fiber.Ctx) error
	GetByIdTodolist(ctx *fiber.Ctx) error
	CreateTodolist(ctx *fiber.Ctx) error
	UpdateTodolist(ctx *fiber.Ctx) error
	DeleteTodolist(ctx *fiber.Ctx) error
}
