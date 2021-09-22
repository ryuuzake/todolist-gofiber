package service

type AuthService interface {
	RegisterUser(email, password string) error
}
