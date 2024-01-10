package authentication

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

type CreateToken struct{}

func (t *CreateToken) CreateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
