package users

import (
	"strings"

	"github.com/psinthorn/gogolang.co/domains/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json: "date_create"`
	Status      bool   `json:"status"`
}

//
// จัดการข้อมูล และตรวจสอบสอบข้อมูลให้ถูกต้อง โดยการสร้าง Method ให้้กับ User Model
//

// Method: Validate
// ตรวจสอบอีเมล์
// จัดการจัดช่องว่างหน้าและหลังอีเมล์
// ทำให้อีเมล์เป็นอักษรตัวเล็ก

func (user *User) Validate() *errors.ErrorRespond {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
