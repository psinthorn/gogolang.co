package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/psinthorn/gogolang.co/services/contents"
)

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome-index.html", nil)
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about-index.html", nil)
}

func AllContents(c *gin.Context) {
	// call GetAllContents on services
	contents, err := services.GetAllContent()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	// if no error then return results to request
	// c.JSON(http.StatusOK, results)
	c.HTML(http.StatusOK, "blog-index.html", contents)
}
