package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/service"
	"github.com/ryuuzake/todolist-gofiber/validator"
	"strconv"
)

type AttachmentControllerImpl struct {
	AttachmentService service.AttachmentService
	FileUploadService service.FileUploadService
}

func (controller AttachmentControllerImpl) GetAllAttachment(ctx *fiber.Ctx) error {
	todolistId, err := strconv.Atoi(ctx.Params("todolistId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	attachments, err := controller.AttachmentService.GetAllAttachmentWithTodolistId(todolistId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(attachments)
}

func (controller AttachmentControllerImpl) GetByIdAttachment(ctx *fiber.Ctx) error {
	attachmentId, err := strconv.Atoi(ctx.Params("attachmentId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	attachment, err := controller.AttachmentService.GetAttachmentById(attachmentId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(attachment)
}

func (controller AttachmentControllerImpl) CreateAttachment(ctx *fiber.Ctx) error {
	todolistId, err := strconv.Atoi(ctx.Params("todolistId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	attachmentPayload := new(validator.AttachmentPayload)
	if err := ctx.BodyParser(attachmentPayload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := validator.ValidateCreateAttachmentPayload(attachmentPayload)
	if validationErr != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(validationErr)
	}

	filePart, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "file not found",
		})
	}

	fileUrl, err := controller.FileUploadService.UploadFile(filePart)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	attachment := model.Attachment{TodolistId: todolistId, Url: fileUrl, Caption: attachmentPayload.Caption}
	if ok := controller.AttachmentService.CreateAttachmentWithTodolistId(todolistId, attachment); ok != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": ok,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func (controller AttachmentControllerImpl) UpdateAttachment(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (controller AttachmentControllerImpl) DeleteAttachment(ctx *fiber.Ctx) error {
	attachmentId, err := strconv.Atoi(ctx.Params("attachmentId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id not recognized",
		})
	}

	if err := controller.AttachmentService.DeleteAttachmentById(attachmentId); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
