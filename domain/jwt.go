package domain

import "github.com/labstack/echo/v4"

type Jwt struct{}

type JwtUsecase interface {
	SetJwtEmployee(g *echo.Group)
	SetJwtAdmin(g *echo.Group)
}
