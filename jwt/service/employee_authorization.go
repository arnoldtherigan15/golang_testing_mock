package service

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *JwtUsecase) SetJwtEmployee(g *echo.Group) {

	secret := os.Getenv("SECRET")
	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(secret),
	}))
	// check if user is employee or not
	g.Use(h.authentication)
	g.Use(h.authorizeEmployee)
}

func (h *JwtUsecase) authorizeEmployee(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// di soal admin bisa akses semua
			if claims["role"] == "employee" || claims["role"] == "admin" {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
