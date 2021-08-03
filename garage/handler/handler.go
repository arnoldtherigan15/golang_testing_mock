package handler

import (
	"net/http"
	"section9/domain"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service domain.GarageService
}

func NewHandler(e *echo.Group, service domain.GarageService) {
	handler := &Handler{service}
	g := e.Group("/garages")
	g.POST("", handler.Create)
	g.GET("", handler.FindAll)
	g.GET("/:id", handler.FindByID)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}

func isRequestValid(garage *domain.Garage) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(garage); err != nil {
		return false, err
	}
	return true, nil
}

func (h *Handler) Create(c echo.Context) (err error) {
	var garage domain.Garage
	if err = c.Bind(&garage); err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&garage); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	createdGarage, err := h.service.Create(&garage)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusCreated, createdGarage)
}

func (h *Handler) FindAll(c echo.Context) (err error) {

	garages, err := h.service.FindAll()

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusCreated, garages)
}

func (h *Handler) Delete(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	isDeleted, err := h.service.Delete(ID)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	response := map[string]bool{
		"is_delete": isDeleted,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Update(c echo.Context) (err error) {
	var garage domain.Garage
	err = c.Bind(&garage)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&garage); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	updatedGarage, err := h.service.Update(ID, &garage)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	response := map[string]bool{
		"is_update": updatedGarage,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) FindByID(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	garage, err := h.service.FindByID(ID)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusOK, garage)
}
