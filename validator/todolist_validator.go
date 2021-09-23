package validator

type CreateTodolistPayload struct {
	Task string `json:"task" xml:"task" form:"task" validate:"required"`
}

type UpdateTodolistPayload struct {
	Task string `json:"task" xml:"task" form:"task" validate:""`
}

func ValidateCreateTodolistPayload(payload *CreateTodolistPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}

func ValidateUpdateTodolistPayload(payload *UpdateTodolistPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}
