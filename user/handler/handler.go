package handler

import (
	"net/http"
	"section9/domain"
	"section9/user/input"
	"section9/user/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service domain.UserService
}

func NewHandler(e *echo.Group, service domain.UserService) {
	handler := &Handler{service}
	g := e.Group("/users")
	g.POST("", handler.Create)
	g.POST("/login", handler.Login)
	g.GET("", handler.FindAll)
	g.GET("/:id", handler.FindByID)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}

func isRequestValid(user *domain.User) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return false, err
	}
	return true, nil
}

func (h *Handler) Create(c echo.Context) (err error) {
	var user domain.User
	if err = c.Bind(&user); err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&user); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	createdUser, err := h.service.Create(&user)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, err)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	formattedUser := response.FormatUserResponse(createdUser)

	return c.JSON(http.StatusCreated, formattedUser)
}

func (h *Handler) Login(c echo.Context) (err error) {
	var user input.LoginInput
	if err = c.Bind(&user); err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	userData, err := h.service.Login(user)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, err)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusOK, userData)
}

func (h *Handler) FindAll(c echo.Context) (err error) {

	users, err := h.service.FindAll()

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}
	formattedUser := response.FormatUsersResponse(users)

	return c.JSON(http.StatusOK, formattedUser)
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
	var user input.UpdateInput
	err = c.Bind(&user)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	updatedUser, err := h.service.Update(ID, &user)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	response := map[string]bool{
		"is_update": updatedUser,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) FindByID(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	user, err := h.service.FindByID(ID)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	formattedUser := response.FormatUserResponse(user)

	return c.JSON(http.StatusOK, formattedUser)
}
