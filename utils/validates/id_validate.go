package utils

import (
	"github.com/psinthorn/gogolang.co/domains/errors"
)

func ValidateId(IdParam string) (id int64, *errors.ErrorRespond) {
	id, err := strconv.ParseInt(IdParam, 10, 64)
	if err != nil {
		idError := errors.NewBadRequestError("user id must be a number")
		return 0, idError
	}
	return id, nil
}
