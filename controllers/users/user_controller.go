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

	result, saveError := services.UserService.CreateUser(user)
	if saveError != nil {

		//Handle Error
		// and return bad request to caller
		c.JSON(saveError.StatusCode, saveError)
		return

	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))

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

	users, getErr := services.UserService.GetAllUser()
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

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

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))

}

func Update(c *gin.Context) {

	// สรุปขั้นตอนการทำงาน
	// 1. ตรวจสอบ method ว่าเป็น patch หรือไม่
	// 2. ตรวจสอบความถูกต้องของ user id ว่าเป็นตัวเลขหรือไม่
	// หากไม่ให้ return error ให้กับ request
	// หากเป็น ไอดี ถูกต้องให้ทำการ
	// 3. ทำการเปรียบเทียบความถูกต้องของข้อมูลที่รับมากับ user model และทำการ Bind ข้อมูลให้เป็น JSON หากข้อมูลไม่ถูกต้องให้ return err หากถูกต้องให้
	// เรียกไปที่ update serivce
	// หมายเหตุ: ใน service จะมีเงื่อนไขในการ validation เพื่อตรวจสอบต่างๆ อีกในขั้นตอนการ update ไปที่ database

	// ตรวจสอบ method ว่าเป็น patch หรือไม่
	// โดยจะคืนค่าเป็น Boolean โดยหากเป็น patch จะคืนค่าเป็น true และเป็น false หาก request method เป็นอื่นๆ
	isPartial := c.Request.Method == http.MethodPatch

	// ตรวจสอบความถูกต้องของ id
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id must be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	// ทำการเปรียบเทียบความถูกต้องของข้อมูลที่รับมากับ user model และทำการ Bind ข้อมูลให้เป็น JSON
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.StatusCode, restError)
		return
	}

	user.Id = userId
	result, updateErr := services.UserService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func Delete(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id must be a number")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		RestErr := errors.NewNotFoundError("user not found")
		c.JSON(RestErr.StatusCode, RestErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

//
// Search User by ID
//

func Search(c *gin.Context) {
	status := c.Query("status")
	// fmt.Sprintf("Status is %s", status)
	// return

	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
