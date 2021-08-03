package handler

import (
	"errors"
	"net/http"
	"section9/domain"
	"section9/order/helpers"
	"section9/order/input"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service domain.OrderService
}

func NewHandler(e *echo.Group, service domain.OrderService) {
	handler := &Handler{service}
	g := e.Group("/order")
	g.POST("", handler.Create)
	g.POST("/done", handler.Done)
	g.GET("", handler.FindAll)
	// g.GET("/:id", handler.FindByID)
	// g.PUT("/:id", handler.Update)
	// g.DELETE("/:id", handler.Delete)
}

func isRequestValid(order *input.Order) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(order); err != nil {
		return false, err
	}
	return true, nil
}

func (h *Handler) Create(c echo.Context) (err error) {
	var order input.Order
	if err = c.Bind(&order); err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}
	if ok, err := isRequestValid(&order); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	createdOrder, err := h.service.Create(&order)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusCreated, createdOrder)
}

func (h *Handler) Done(c echo.Context) (err error) {
	var order input.DoneInput
	if err = c.Bind(&order); err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	doneOrder, err := h.service.Done(&order)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusOK, doneOrder)
}

func (h *Handler) FindAll(c echo.Context) (err error) {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, errors.New("page is required"))
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, errors.New("limit is required"))
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	order := c.QueryParam("order")
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, errors.New("limit is required"))
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	pagination := helpers.Pagination{
		Page:  page,
		Limit: limit,
		Sort:  order,
	}
	data, err := h.service.FindAll(pagination)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	type Meta struct {
		TotalRows  int64 `json:"total_data"`
		TotalPages int   `json:"total_pages"`
		Limit      int   `json:"limit"`
	}

	meta := Meta{data.TotalRows, data.TotalPages, data.Limit}

	formattedData := map[string]interface{}{
		"meta": meta,
		"rows": data.Rows,
	}

	return c.JSON(http.StatusCreated, formattedData)
}
