package contents

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/domains/contents"
	"github.com/psinthorn/gogolang.co/domains/errors"
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
		fmt.Printf(jsonErr.Error)
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
	results, err := services.GetAllContent()
	if err != nil {
		c.JSON(err.StatusCode, err)
	}
	c.JSON(http.StatusOK, results)
}

// GET content by ID
func Get(c *gin.Context) {

	//STEP:
	// - check ID is int64 and valid
	// - call service to get content by given id
	// - if no contetn found by given id return notfond error
	// - or return content to request

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		restErr := errors.NewBadRequestError("content id must be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	content, getErr := services.GetContent(id)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, content)

}

// Update content by ID
func Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me please")
}

// delete content by ID
func Delete(c *gin.Context) {
	id, err := utils.ValidateId(c.Param("id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	_, err = services.CreateContent(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]string{"content": "deleted"})
}
