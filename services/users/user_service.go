package services

import (
	"strings"

	models "github.com/psinthorn/gogolang.co/domain/errors"
	"github.com/psinthorn/gogolang.co/domain/users"
)

func CreateUser(user users.User) (*users.User, *models.ErrorRespond) {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	return &user, nil
}
