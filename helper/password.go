package helper

import "golang.org/x/crypto/bcrypt"

func GenerateHashFromPassword(password string) (string, error) {
	// TODO: Add BCrypt Cost to Config?
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 5)

	return string(hashedPassword), err
}

func CompareHashAndPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
