package contents

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/psinthorn/F2Go/services/contents"
	"github.com/psinthorn/gogolang.co/domain/contents"
	"github.com/psinthorn/gogolang.co/domain/errors"
)

//
// CRUD operation controllers
//

func Create(c *gin.Context) {
	var content contents.Content

	if err := c.ShouldBindJSON(&content); err != nil {
		jsonErr := errors.NewBadRequestError("invalid Json body")
		c.JSON(jsonErr.StatusCode, jsonErr)
		return
	}

	resultContent, saveErr := services.CreateContent(content)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}

	c.JSON(http.StatusOK, resultContent)
}

func Get(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me please")
}

func Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me please")
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me please")
}
