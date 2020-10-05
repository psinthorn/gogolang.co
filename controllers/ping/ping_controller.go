package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ping struct{}

var Ping ping

func (p *ping) Pong(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
