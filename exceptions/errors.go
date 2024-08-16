package exceptions

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

func InternalError() *CustomError {

	return NewCustomError(http.StatusInternalServerError, "Internal server error")

}

func NotFound() *CustomError {

	return NewCustomError(http.StatusNotFound, "Resource could not be found")

}

func BadRequest() *CustomError {

	return NewCustomError(http.StatusBadRequest, "Bad Request")

}
