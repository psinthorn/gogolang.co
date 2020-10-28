package services

import (
	"fmt"

	"github.com/psinthorn/gogolang.co/domains/errors"
	"github.com/psinthorn/gogolang.co/domains/users"
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

func UpdateUser(user users.User) (*users.User, *errors.ErrorRespond) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println("befor update: ", currentUser)

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email
	currentUser.Avatar = user.Avatar
	currentUser.Status = user.Status

	if err := currentUser.Update(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("after updated: ", currentUser)

	return currentUser, nil
}

func DeleteUser(user users.User) (*users.User, *errors.ErrorRespond) {
	return nil, nil
}
