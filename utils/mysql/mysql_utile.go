package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/psinthorn/gogolang.co/domains/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func PareError(err error) *errors.ErrorRespond {
	mysqlErr, ok := err.(*mysql.MySQLError)
	fmt.Println(err.Error())
	fmt.Println(ok)

	if !ok {
		if strings.Contains("no rows in result set", errorNoRows) {
			return errors.NewNotFoundError("no record found with given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch mysqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data or duplicate data")
	}
	return errors.NewInternalServerError("error process request")
}
