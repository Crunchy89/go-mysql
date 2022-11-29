package s

import (
	"net/http"

	"github.com/Crunchy89/go-mysql/utils/r"
)

type Result struct {
	Code         int         `json:"code"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func NewResultData(data interface{}) *Result {
	return &Result{
		Code:         http.StatusOK,
		Data:         data,
		ErrorMessage: "",
	}
}

func NewResultMessage(message string) *Result {
	return &Result{
		Code:         http.StatusOK,
		Data:         message,
		ErrorMessage: "",
	}
}

func NewResultError(err error) *Result {
	code := http.StatusInternalServerError
	if eR, ok := err.(r.Ex); ok {
		if eR.IsDataNotFound() {
			code = 404
		} else if eR.IsDatabaseError() {
			code = 500
		} else if eR.IsRepositoryError() {
			code = 500
		} else if eR.IsServiceError() {
			code = 500
		}
	}
	return &Result{
		Code:         code,
		Data:         nil,
		ErrorMessage: err.Error(),
	}
}
