package services

import (
	"fmt"

	"github.com/psinthorn/gogolang.co/domains/errors"
	"github.com/psinthorn/gogolang.co/domains/users"
	"github.com/psinthorn/gogolang.co/utils/crypto_utils"
	date_utils "github.com/psinthorn/gogolang.co/utils/date"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.ErrorRespond)
	GetAllUser() (users.Users, *errors.ErrorRespond)
	GetUser(int64) (*users.User, *errors.ErrorRespond)
	UpdateUser(bool, users.User) (*users.User, *errors.ErrorRespond)
	DeleteUser(int64) *errors.ErrorRespond
	SearchUser(string) (users.Users, *errors.ErrorRespond)
}

func (u *userService) CreateUser(user users.User) (*users.User, *errors.ErrorRespond) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	// set user date created
	// use standard time zone
	// for Thailand need to +7 from UTC
	// now := time.Now().UTC()
	// user.DateCreated = now.Format("2006-01-02T15:04:05Z")
	user.DateCreated = date_utils.GetNowDbDateLayout()
	user.Password = crypto_utils.Md5Encrypt(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userService) GetAllUser() (users.Users, *errors.ErrorRespond) {

	dao := &users.User{}
	return dao.GetAll()
}

func (u *userService) GetUser(userId int64) (*users.User, *errors.ErrorRespond) {
	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.ErrorRespond) {
	currentUser, err := UserService.GetUser(user.Id)
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

func (u *userService) DeleteUser(userId int64) *errors.ErrorRespond {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (u *userService) SearchUser(status string) (users.Users, *errors.ErrorRespond) {
	dao := &users.User{}
	return dao.FindUserByStatus(status)

}
