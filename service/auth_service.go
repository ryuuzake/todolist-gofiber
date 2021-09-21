package service

type AuthService interface {
	RegisterUser(email, password string) (interface{}, error)
}
