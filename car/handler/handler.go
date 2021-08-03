package handler

import (
	"net/http"

	// "os"
	"section9/domain"
	"strconv"

	"github.com/go-playground/validator/v10"
	// mid "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	CarService domain.CarService
}

func NewHandler(e *echo.Group, carService domain.CarService) {
	handler := &Handler{carService}
	g := e.Group("/cars")
	g.POST("", handler.Create)
	g.GET("", handler.FindAll)
	g.GET("/:id", handler.FindByID)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
	g.PATCH("/:id/garage", handler.UpdateGarage)
}

func isRequestValid(car *domain.Car) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(car); err != nil {
		return false, err
	}
	return true, nil
}

func (h *Handler) FindAll(c echo.Context) (err error) {
	cars, err := h.CarService.FindAll()

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusOK, cars)
}

func (h *Handler) Delete(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	isDeleted, err := h.CarService.Delete(ID)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	response := map[string]bool{
		"is_delete": isDeleted,
	}

	return c.JSON(http.StatusOK, response)
}
func (h *Handler) Create(c echo.Context) (err error) {
	var car domain.Car
	err = c.Bind(&car)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&car); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	createdCar, err := h.CarService.Create(&car)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusCreated, createdCar)
}

func (h *Handler) Update(c echo.Context) (err error) {
	var car domain.Car
	err = c.Bind(&car)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&car); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	updatedCar, err := h.CarService.Update(ID, &car)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	response := map[string]bool{
		"is_update": updatedCar,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) FindByID(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	car, err := h.CarService.FindByID(ID)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusOK, car)
}

func (h *Handler) UpdateGarage(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	garageID, err := strconv.Atoi(c.QueryParam("garage_id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	updatedCar, err := h.CarService.UpdateGarage(ID, garageID)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	response := map[string]bool{
		"is_update_garage": updatedCar,
	}

	return c.JSON(http.StatusOK, response)
}
