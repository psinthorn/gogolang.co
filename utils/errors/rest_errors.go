package errors

import (
	"net/http"

	models "github.com/psinthorn/gogolang.co/domain/errors"
)

//
// Error model type locate at domain/errors/error_dto.go
//

func NewBadRequestError(message string) *models.ErrorRespond {

	return &models.ErrorRespond{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}
