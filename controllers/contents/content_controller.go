package contents

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/domains/contents"
	"github.com/psinthorn/gogolang.co/domains/errors"
	services "github.com/psinthorn/gogolang.co/services/contents"
	validate_utils "github.com/psinthorn/gogolang.co/utils/validates"
)

//
// CRUD operation controllers
// standard content
//

var content contents.Content

// var isApi bool = false

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
	// // Check Application Type if JSON will be true
	isApi, err := validate_utils.IsApi(c.Param("api"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	// call GetAllContents on services
	contents, err := services.GetAllContent()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	if isApi {
		c.JSON(http.StatusOK, contents)
	} else {
		c.HTML(http.StatusOK, "blog-index.html", contents)
	}

	// // if no error then return results to request
	// if api {
	// 	c.JSON(http.StatusOK, contents)
	// } else {
	// 	c.HTML(http.StatusOK, "blog-index.html", contents)
	// }

}

// GET content by ID
func Get(c *gin.Context) {

	//STEP:
	// - check ID is int64 and valid
	// - call service to get content by given id
	// - if no contetn found by given id return notfond error
	// - or return content to request

	isApi, err := validate_utils.IsApi(c.Param("api"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	id, err := validate_utils.Id(c.Param("id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	content, getErr := services.GetContent(id)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	if isApi {
		c.JSON(http.StatusOK, content)
	} else {
		c.HTML(http.StatusOK, "blog-show.html", content)
	}

}

// Update content by ID
func Update(c *gin.Context) {
	// Validate ID
	// Validate method PUT or PATCH
	// ShouldBindJson
	// If Error or can't bind to Json return error to request
	// If can bund json Call ContentUpdat Method on Content Service
	// If Error return error to request
	// if Update success return rsult to request

	// id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// if err != nil {
	// 	restErr := errors.NewBadRequestError("content id must be a number")
	// 	c.JSON(restErr.StatusCode, restErr)
	// 	return
	// }

	fmt.Println(c.Param("api"))
	// 	isApi, err := validate_utils.IsApi(c.Param("api")) {
	// 	if err != nil {
	// 		c.JSON(err.StatusCode, err)
	// 		return
	// 	}
	// }

	// isApi, err := strconv.ParseInt(c.Param("api"), 10, 64)
	// if err != nil {
	// 	restErr := errors.NewBadRequestError("content id must be a number")
	// 	c.JSON(restErr.StatusCode, restErr)
	// 	return
	// }

	// isPartial := c.Request.Method == http.MethodPatch

	// if isApi {
	// 	c.JSON(http.StatusNotImplemented, "Implement me please")
	// } else {
	// 	c.HTML(http.StatusOK, "update-form.html", nil)
	// }

}

// delete content by ID
func Delete(c *gin.Context) {
	// validate id
	// call DeleteContent services
	// if err not nil return err
	// if success return c.JSON with status ok and map[string]string{"status": "deleted"} to request
	Id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("content id must be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	if err := services.DeleteContent(Id); err != nil {
		RestErr := errors.NewNotFoundError("content not found")
		c.JSON(RestErr.StatusCode, RestErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}
