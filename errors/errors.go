package errors

import (
	"fmt"
	"net/http"
)

type LogError struct {
	Msg   string
	Args  string
	Error error
}

type CustomError struct {
	Code    int    // HTTP status code
	Message string // Error message
}

// Error method to satisfy the error interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewCustomError(code int, msg string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: msg,
	}
}

func InternalError(code int, msg string) *CustomError {
	if msg == "" {
		return NewCustomError(http.StatusInternalServerError, "Internal server error")
	}
	return NewCustomError(http.StatusInternalServerError, msg)
}

func NotFound(code int, msg string) *CustomError {
	if msg == "" {
		return NewCustomError(http.StatusNotFound, "Resource could not be found")
	}
	return NewCustomError(http.StatusNotFound, msg)
}

func BadRequest(code int, msg string) *CustomError {
	if msg == "" {
		return NewCustomError(http.StatusBadRequest, "Bad Request")
	}
	return NewCustomError(http.StatusBadRequest, msg)
}
