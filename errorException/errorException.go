package errorException

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusNotFound}
}

func NewUexpectedError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusInternalServerError}
}
