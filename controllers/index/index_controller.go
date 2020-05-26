package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type indexPage struct{}

var IndexPage indexPage

func (i *indexPage) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome-index.html", nil)
}

func (i *indexPage) About(c *gin.Context) {
	c.HTML(http.StatusOK, "about-index.html", nil)
}
