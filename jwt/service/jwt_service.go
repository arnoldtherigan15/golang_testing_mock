package service

import (
	"fmt"
	"net/http"
	"section9/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type JwtUsecase struct {
	UserRepo domain.UserRepository
}

func NewJwtService(userRepo domain.UserRepository) domain.JwtUsecase {
	return &JwtUsecase{userRepo}
}

func (h *JwtUsecase) authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			email := fmt.Sprintf("%v", claims["email"])
			userData, err := h.UserRepo.FindByEmail(email)
			if err != nil || userData.ID == 0 {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}
			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
