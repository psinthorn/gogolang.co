package services

import "github.com/psinthorn/gogolang.co/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
