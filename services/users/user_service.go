package services

import (
	"fmt"

	"github.com/psinthorn/gogolang.co/domains/errors"
	"github.com/psinthorn/gogolang.co/domains/users"
	date_utils "github.com/psinthorn/gogolang.co/utils/date"
)

func CreateUser(user users.User) (*users.User, *errors.ErrorRespond) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	// set user date created
	// use standard time zone
	// for Thailand need to +7 from UTC
	// now := time.Now().UTC()
	// user.DateCreated = now.Format("2006-01-02T15:04:05Z")
	user.DateCreated = date_utils.GetNowDbDateLayout()
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

	if isPartial {
		fmt.Println(isPartial)
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}

		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}

		if user.Email != "" {
			currentUser.Email = user.Email
		}

		if user.Avatar != "" {
			currentUser.Avatar = user.Avatar
		}

		if user.Status != "" {
			currentUser.Status = user.Status
		}

	} else {

		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
		currentUser.Avatar = user.Avatar
		currentUser.Status = user.Status

	}

	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }

	if err := currentUser.Update(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	// updateUser, err := GetUser(user.Id)
	// if err != nil {
	// 	return nil, err
	// }

	return currentUser, nil
}

func DeleteUser(userId int64) *errors.ErrorRespond {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.ErrorRespond) {

	return users.FinduserByStatus(status)

}
