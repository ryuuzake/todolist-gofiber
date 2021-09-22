package validator

import "strings"

type RegisterUserPayload struct {
	Email                string `json:"email" xml:"email" form:"email" validate:"required,email,min=6,max=32"`
	Password             string `json:"password" xml:"password" form:"password" validate:"required,min=8,max=32"`
	PasswordConfirmation string `json:"password_confirmation" xml:"password_confirmation" form:"password_confirmation" validate:"required,min=8,max=32"`
}

type LoginUserPayload struct {
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
}

func ValidateRegisterUserPayload(payload *RegisterUserPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	// Additional Validation
	// Check password
	if strings.Compare(payload.Password, payload.PasswordConfirmation) != 0 {
		element := MakeErrorResponse("PasswordConfirmation", "same-as-password", "")
		errors = append(errors, &element)
	}

	return errors
}

func ValidateLoginUserPayload(payload *LoginUserPayload) []*ErrorResponse {
	errors := ValidateStruct(*payload)

	return errors
}
