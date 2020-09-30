package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome-index.html", nil)
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about-index.html", nil)
}
