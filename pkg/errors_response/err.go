package errors_response

import "net/http"

type ErrorResponse interface {
	Message() string
	Status() int
	Error() string
}

type ErrorData struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *ErrorData) Message() string {
	return e.ErrMessage
}

func (e *ErrorData) Status() int {
	return e.ErrStatus
}

func (e *ErrorData) Error() string {
	return e.ErrError
}

func NewInternalServerError(message string) ErrorResponse {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewNotFoundError(message string) ErrorResponse {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "NOT_FOUND",
	}
}

func NewBadRequestResponse(message string) ErrorResponse {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus: http.StatusBadRequest,
		ErrError: "BAD_REQUEST",
	}
}

func NewUnprocessableEntityResponse(message string) ErrorResponse {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus: http.StatusUnprocessableEntity,
		ErrError: "INVALID_REQUEST_BODY",
	}
}