package services

import (
	domain "github.com/psinthorn/gogolang.co/domain/users"
	domain "github.com/psinthorn/gogolang.co/domain/utils"
)

type users struct{}

var Users users

func (u *users) CreateUser() (*domain.User, *domain.AppErr) {

}

func (u *users) GetUser(userId int64) (*domain.User, *domain.AppErr) {

}
