package app_error

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    // HTTP Status Code
	Err     error  // Error message
	Context string // Which file/Method occured the error
}

func (e *AppError) StatusCode() int {
	return e.Code
}

// Return error instance
func (e *AppError) Error() error {
	return e.Err
}

func (e *AppError) Message() string {
	return e.Err.Error()
}

// Message with Context
func (e *AppError) MessageContext() string {
	return fmt.Sprintf("%s: %s", e.Context, e.Message())
}

func NewNotFoundError(message, context string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Err:     fmt.Errorf(message),
		Context: context,
	}
}

func NewUnexpectedError(message, context string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Err:     fmt.Errorf(message),
		Context: context,
	}
}
