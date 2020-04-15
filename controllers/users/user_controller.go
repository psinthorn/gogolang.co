package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/psinthorn/gogolang.co/services/users"
)

type users struct{}

var Users users

func (u *users) CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}

func (u *users) GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
	services.Users.GetUser(user_id)
}

func (u *users) SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}

func (u *users) UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}

func (u *users) DelleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}
