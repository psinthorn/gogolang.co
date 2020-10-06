package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/domain/users"
	services "github.com/psinthorn/gogolang.co/services/users"
	
)

//
// Create new user
//
func Create(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("In valid Json")
		c.JSON(restError.StatusCode, restError)
		return
	}

	if err := user.Validate(); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {

		//Handle Error
		// and return bad request to caller
		c.JSON(saveError.StatusCode, saveError)
		return

	}
	c.JSON(http.StatusCreated, result)

	// // Traditional How to read data from request input
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO: handle error if err not nil
	// 	return
	// }

	// Traditional how to bind data to data model
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	//TODO: handle error if err not nil
	// 	return
	// }

}

//
// Get user by ID
//
func Get(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id must be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}

func Update(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}

func Delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}
