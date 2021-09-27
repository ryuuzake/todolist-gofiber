package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ryuuzake/todolist-gofiber/config"
	"github.com/ryuuzake/todolist-gofiber/model"
	"time"
)

type UserClaim struct {
	UserId       int
	UserFullName string
	RoleName     string
	Expire       time.Time
}

func DecodeJWT(ctx *fiber.Ctx) UserClaim {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userIdClaim := 0
	if claims["sub"] != nil {
		userIdClaim = int(claims["sub"].(float64))
	}

	userNameClaim := ""
	if claims["name"] != nil {
		userNameClaim = claims["name"].(string)
	}

	roleClaim := ""
	if claims["role"] != nil {
		roleClaim = claims["role"].(string)
	}

	expireClaim := time.Now()
	if claims["exp"] != nil {
		expireClaim = time.Unix(int64(claims["exp"].(float64)), 0)
	}

	userClaim := UserClaim{
		UserId:       userIdClaim,
		UserFullName: userNameClaim,
		RoleName:     roleClaim,
		Expire:       expireClaim,
	}

	return userClaim
}

func EncodeJWT(user model.User) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Id
	claims["name"] = user.FullName

	// Claim for role
	claims["role"] = user.Role.Name

	// 72 hours = 3 days, from now
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString(config.Config.JWTSecret)
}
