package users

import (
	"fmt"

	"github.com/psinthorn/gogolang.co/domain/errors"
	
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.ErrorRespond {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user_id %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	user.Status = result.Status

	return nil
}

func (user *User) Save() *errors.ErrorRespond {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email address %s aleady registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user id %d is already exist", user.Id))
	}

	usersDB[user.Id] = user
	return nil
}
