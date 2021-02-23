package users

import (
	"strings"

	"github.com/psinthorn/gogolang.co/domains/errors"
)

type User struct {
	Id          int64  `json:"id" bson:"id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Email       string `json:"email" bson:"email"`
	Avatar      string `json: "avatar" bson: "avatar"`
	Password    string `json: "password" bson: "password"`
	Status      string `json:"status" bson:"status"`
	DateCreated string `json: "date_created" bson: "date_created"`
	// DateUpdated string `json: "date_created" bson: "date_created"`
}

type Users []User

//
// จัดการข้อมูล และตรวจสอบสอบข้อมูลให้ถูกต้อง โดยการสร้าง Method ให้้กับ User Model
//

// Method: Validate
// ตรวจสอบอีเมล์
// จัดการจัดช่องว่างหน้าและหลังอีเมล์
// ทำให้อีเมล์เป็นอักษรตัวเล็ก

// Validate จะทำการตรวจสอบความถูกต้องของข้อมูล
func (user *User) Validate() *errors.ErrorRespond {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Password can not be an empty")
	}

	return nil
}
