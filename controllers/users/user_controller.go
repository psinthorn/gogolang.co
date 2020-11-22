package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/gogolang.co/domains/errors"
	"github.com/psinthorn/gogolang.co/domains/users"
	services "github.com/psinthorn/gogolang.co/services/users"
)

//
// Create new user
//
func Create(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
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
// Get all users
//
func GetAll(c *gin.Context) {

	users, getErr := services.GetAllUser()
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, users)

}

//
// Get user by ID
//
func Get(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
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

func Update(c *gin.Context) {
	// ตรวจสอบความถูกต้องของ user id ว่าเป็นตัวเลขหรือไม่
	// หากไม่ให้ return error ให้กับ request

	// หากเป็น ไอดี ถูกต้องให้ทำการ
	// เรียกไปที่ update serivce
	// หมายเหตุ: ใน service จะมีเงื่อนไขในการ validation เพื่อตรวจสอบต่างๆ อีกในขั้นตอนการ update ไปที่ database
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id must be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.StatusCode, restError)
		return
	}

	user.Id = userId
	result, updateErr := services.UpdateUser(user)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func Delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me Please")
}
