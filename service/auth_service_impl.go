package service

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ryuuzake/todolist-gofiber/model"
	"github.com/ryuuzake/todolist-gofiber/repository"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func (service *AuthServiceImpl) RegisterUser(email, password string) error {
	if _, err := service.UserRepository.FindByEmail(email); err == nil {
		return errors.New("user already registered")
	}

	err := service.UserRepository.Create(model.User{Email: email, Password: password})
	if err != nil {
		return err
	}

	return nil
}

func (service *AuthServiceImpl) LoginUser(email string, password string) (string, error) {
	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	role, err := service.RoleRepository.FindById(user.RoleId)
	if err != nil {
		return "", errors.New("role not found")
	}

	user.Role = role

	// TODO: Check User password hash
	if strings.Compare(user.Password, password) != 0 {
		return "", errors.New("user creds failed")
	}

	return service.generateJWTToken(user)
}

func (service *AuthServiceImpl) generateJWTToken(user model.User) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Id
	claims["name"] = user.FullName

	// Claim for role
	claims["role"] = user.Role.Name

	// TODO: Make Expire Time more readable
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	// TODO: Random Generated Key same as in auth_middleware.go
	return token.SignedString([]byte("secret"))
}
