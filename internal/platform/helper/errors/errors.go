package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	message string
	error   string
	cause   string
	Status  int
}

func (e ApiError) Error() string {
	return e.message
}
func (e ApiError) GetStatus() int {
	return e.Status
}

func (e ApiError) toMap() map[string]interface{} {
	return map[string]interface{}{
		"message": e.message,
		"errors":  e.error,
		"status":  e.Status,
		"cause":   e.cause,
	}
}

func New(message string) error {
	return errors.New(message)
}

func StatusUnauthorizedApiError(message string, cause string) error {
	return ApiError{message, "unauthorized_request", cause, http.StatusUnauthorized}
}

func InternalServerApiError(message string, cause string) error {
	return ApiError{message, "internal server error", cause, http.StatusInternalServerError}
}

func ForbiddenApiError(message string, cause string) error {
	return ApiError{message, "forbidden_resource", cause, http.StatusForbidden}
}

func NotFoundApiError(message string, cause string) error {
	return ApiError{message, "not_found", cause, http.StatusNotFound}
}
func BadRequestApiError(message string, cause string) error {
	return ApiError{message, "bad_request", cause, http.StatusBadRequest}
}

func ConflictApiError(message string, cause string) error {
	return ApiError{message, "bad_request", cause, http.StatusConflict}
}

func RespondError(c *gin.Context, e error) {
	if err, ok := e.(ApiError); ok {
		c.JSON(err.Status, e.(ApiError).toMap())
	} else {
		errDefault := New(e.Error())
		c.JSON(http.StatusInternalServerError, errDefault)
	}
}
