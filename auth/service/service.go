package service

import (
	"errors"
	"fmt"
	"os"
	"section9/domain"

	"github.com/dgrijalva/jwt-go"
)

type service struct{}

func NewService() *service {
	return &service{}
}

type Service interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

func (s *service) GenerateToken(user *domain.User) (string, error) {
	claim := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	var SECRET_KEY = os.Getenv("SECRET")

	access_token, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return access_token, err
	}

	return access_token, nil
}

func (s *service) ValidateToken(encodedToken string) (*jwt.Token, error) {
	access_token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		var SECRET_KEY = os.Getenv("SECRET")
		fmt.Println()
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return access_token, err
	}
	return access_token, nil
}
