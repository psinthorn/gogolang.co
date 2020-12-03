package validate_utils

import (
	"strconv"

	"github.com/psinthorn/gogolang.co/domains/errors"
)

func Id(IdParam string) (int64, *errors.ErrorRespond) {
	id, err := strconv.ParseInt(IdParam, 10, 64)
	if err != nil {
		idError := errors.NewBadRequestError("user id must be a number")
		return 0, idError
	}
	return id, nil
}

func IsApi(api string) (bool, *errors.ErrorRespond) {
	isApi, err := strconv.ParseBool(api)
	if err != nil {
		apiError := errors.NewBadRequestError("api must be a boolean")
		return false, apiError
	}
	return isApi, nil
}
