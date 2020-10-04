package services

import (
	models "github.com/psinthorn/gogolang.co/domain/errors"
	"github.com/psinthorn/gogolang.co/domain/users"
)

func CreateUser(user users.User) (*users.User, *models.ErrorRespond) {
	return &user, nil
}
