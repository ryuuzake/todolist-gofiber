package validator

import "github.com/gofrs/uuid"

type CreateTodolistPayload struct {
	Task string `json:"task" xml:"task" form:"task" validate:"required"`
}

type UpdateTodolistPayload struct {
	Task string `json:"task" xml:"task" form:"task" validate:""`
}

type TodoPayload struct {
	Title string `json:"title" xml:"title" form:"title" validate:"required"`
}

type TodolistPayload struct {
	Task     string    `json:"task" validate:"required"`
	StatusId uuid.UUID `json:"status_id"`
}

type AttachmentPayload struct {
	Caption string `form:"caption" validate:"required"`
}

func ValidateCreateTodoPayload(payload *TodoPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}

func ValidateUpdateTodoPayload(payload *TodoPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}

func ValidateCreateTodolistPayload(payload *TodolistPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}

func ValidateUpdateTodolistPayload(payload *TodolistPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}

func ValidateCreateAttachmentPayload(payload *AttachmentPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}
