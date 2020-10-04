package users

import (
	"strings"

	models "github.com/psinthorn/gogolang.co/domain/errors"
	"github.com/psinthorn/gogolang.co/utils/errors"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    bool   `json:"status"`
}

//
// จัดการข้อมูล และตรวจสอบสอบข้อมูลให้ถูกต้อง โดยการสร้าง Method ให้้กับ User Model
//

// Method: Validate
// ตรวจสอบอีเมล์
// จัดการจัดช่องว่างหน้าและหลังอีเมล์
// ทำให้อีเมล์เป็นอักษรตัวเล็ก

func (user *User) Validate() *models.ErrorRespond {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		err := errors.NewBadRequestError("Invalid Email Address")
		return err
	}
	return nil
}
