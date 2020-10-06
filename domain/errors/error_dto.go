package errors

import "net/http"

type ErrorRespond struct {
	//Id         string `json: "id"`
	Message    string `json: "message"`
	StatusCode int    `json: "status_code"`
	Error      string `json: "error"`
}

//
// Error model type locate at domain/errors/error_dto.go
//

func NewBadRequestError(message string) *ErrorRespond {
	return &errors.ErrorRespond{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

func NewNotFoundError(message string) *ErrorRespond {
	return &errors.ErrorRespond{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

func NewContentAlertNotice(message string) *ErrorRespond {
	return &errors.ErrorRespond{
		Message:    message,
		StatusCode: http.StatusNoContent,
		Error:      "no_content",
	}
}
