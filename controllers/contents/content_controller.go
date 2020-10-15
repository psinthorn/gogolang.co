package contents

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/domain/contents"
	"github.com/psinthorn/gogolang.co/domain/errors"
	services "github.com/psinthorn/gogolang.co/services/contents"
)

//
// CRUD operation controllers
// standard content
//

var content contents.Content

func Create(c *gin.Context) {
	var content contents.Content

	// bind data request to content model as josn
	// if request with invalid data will turn an error
	if err := c.ShouldBindJSON(&content); err != nil {
		jsonErr := errors.NewBadRequestError("invalid Json body")
		c.JSON(jsonErr.StatusCode, jsonErr)
		return
	}

	// create data on database by calling service
	// if error not nil return error to user
	resultContent, saveErr := services.CreateContent(content)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}

	// create succuess return created content to user
	c.JSON(http.StatusOK, resultContent)
}

// GET content all
func GetAll(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement Get all contents please")
}

// GET content by ID
func Get(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me please")
}

// Update content by ID
func UpdateContent(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me please")
}

// delete content by ID
func DeleteContent(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me please")
}
