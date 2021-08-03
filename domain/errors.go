package domain

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")
	ErrBadRequest    = errors.New("given Request is not valid")
)

type Error struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func ErrorFormatter(status int, err error) Error {
	return Error{status, err.Error()}
}

func ErrorValidationFormatter(err error) (msg error) {
	temp := ""
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				if temp != "" {
					temp += ", "
				}
				temp += fmt.Sprintf("%s is required",
					err.Field())

				msg = fmt.Errorf("%s", temp)
			}
		}
	}
	return
}
