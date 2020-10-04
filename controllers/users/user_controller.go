package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/domain/users"
	services "github.com/psinthorn/gogolang.co/services/users"
	utils "github.com/psinthorn/gogolang.co/utils/errors"
)

func Create(c *gin.Context) {

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := utils.NewBadRequestError("In valid Json")
		c.JSON(restError.StatusCode, restError)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {

		//Handle Error
		// and return bad request to caller
		c.JSON(saveErr.StatusCode, saveErr)
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

func Get(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")

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
