package validator

type CreateTodolistPayload struct {
	Task string `json:"task" xml:"task" form:"task" validate:"required"`
}

type UpdateTodolistPayload struct {
	Task string `json:"task" xml:"task" form:"task" validate:""`
}

type TodoPayload struct {
	Title string `json:"title" xml:"title" form:"title" validate:"required"`
}

func ValidateCreateTodoPayload(payload *TodoPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}

func ValidateUpdateTodoPayload(payload *TodoPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}
