package utils

import (
	"net/http"

	models "github.com/psinthorn/gogolang.co/domain/errors"
)

// type RestError struct {
// 	Id         string `json: "id"`
// 	Message    string `json: "message"`
// 	StatusCode int    `json: "status_code"`
// 	Error      string `json: "error"`
// }

func NewBadRequestError(message string) *models.ErrorRespond {

	return &models.ErrorRespond{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}
