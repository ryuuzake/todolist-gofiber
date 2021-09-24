package service

type AuthService interface {
	RegisterUser(email, password string) error
	LoginUser(email, password string) (string, error)
}
