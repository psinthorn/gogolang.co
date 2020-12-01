package services

import (
	"fmt"

	"github.com/psinthorn/gogolang.co/domains/errors"
	"github.com/psinthorn/gogolang.co/domains/users"
)

func CreateUser(user users.User) (*users.User, *errors.ErrorRespond) {

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUser() ([]users.User, *errors.ErrorRespond) {

	allUsers, err := users.GetAll()
	if err != nil {
		return nil, err
	}
	return allUsers, nil
}

func GetUser(userId int64) (*users.User, *errors.ErrorRespond) {
	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.ErrorRespond) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	// if isPartial {
	// 	fmt.Println(isPartial)
	// 	if user.FirstName == "" {
	// 		user.FirstName = currentUser.FirstName
	// 	}

	// 	if user.LastName == "" {
	// 		user.LastName = currentUser.LastName
	// 	}

	// 	if user.Email == "" {
	// 		user.Email = currentUser.Email
	// 	}

	// 	if user.Avatar == "" {
	// 		user.Avatar = currentUser.Avatar
	// 	}

	// 	return currentUser, nil

	// }
	fmt.Println("new user data: ", user)
	fmt.Println("befor update: ", currentUser)
	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email
	currentUser.Avatar = user.Avatar
	currentUser.Status = user.Status
	fmt.Println("after update: ", currentUser)

	if err := currentUser.Update(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	updateUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func DeleteUser(userId int64) *errors.ErrorRespond {
	user := &users.User{Id: userId}
	return user.Delete()
}
