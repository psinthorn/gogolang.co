// package errors

// import (
//	"net/http"

//	"github.com/psinthorn/gogolang.co/domain/errors"
//)

//
// Error model type locate at domain/errors/error_dto.go
//

//func NewBadRequestError(message string) *errors.ErrorRespond {
//	return &errors.ErrorRespond{
//		Message:    message,
//		StatusCode: http.StatusBadRequest,
//		Error:      "bad_request",
//	}
// }

//func NewNotFoundError(message string) *errors.ErrorRespond {
//	return &errors.ErrorRespond{
//		Message:    message,
//		StatusCode: http.StatusNotFound,
//		Error:      "not_found",
//	}
// }

// func NewContentAlertNotice(message string) *errors.ErrorRespond {
//	return &errors.ErrorRespond{
//		Message:    message,
//		StatusCode: http.StatusNoContent,
//		Error:      "no_content",
//	}
//}
