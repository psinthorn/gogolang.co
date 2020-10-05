package services

import (
	"github.com/psinthorn/gogolang.co/domain/errors"
	"github.com/psinthorn/gogolang.co/domain/users"
)

func GetUser(userId int64) (*users.User, *errors.ErrorRespond) {
	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.ErrorRespond) {

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
