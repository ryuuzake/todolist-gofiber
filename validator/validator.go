package validator

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(obj interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()

	err := validate.Struct(obj)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := MakeErrorResponse(err.StructNamespace(), err.Tag(), err.Param())

			errors = append(errors, &element)
		}
	}

	return errors
}

func MakeErrorResponse(field, tag, value string) ErrorResponse {
	return ErrorResponse{FailedField: field, Tag: tag, Value: value}
}
