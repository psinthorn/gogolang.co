package users

import (
	"fmt"
	"strings"

	mysql_db "github.com/psinthorn/gogolang.co/datasources/mysql/users_db"
	"github.com/psinthorn/gogolang.co/domains/errors"
	date_utils "github.com/psinthorn/gogolang.co/utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO USERS(first_name, last_name, email, avatar, status, date_created) VALUES(?,?,?,?,?,?);"
	queryGetAllUsers = "SELECT * FROM users"
	queryGetUserById = "SELECT id, first_name, last_name, email, avatar, status, date_created FROM users WHERE id = ?"
	errorNoRows      = "no rows in result set"
)

var (
	usersDB = make(map[int64]*User)
)

// Get user by id from database
func (user *User) Get() *errors.ErrorRespond {

	// ใช้ Prepare เพื่อตรวจสอบความถูกต้องของข้อมูลก่อนที่จะส่งไปทำการ  process ที่ server เพื่อลดการทำงาน process ที่ฝั่ง server
	stmt, err := mysql_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Avatar, &user.DateCreated, &user.Status); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewBadRequestError(fmt.Sprintf("user id: %d is not found", user.Id))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error on trying to get user id: %d errors: %s ", user.Id, err.Error()))
	}

	return nil
}

// Save user to database
func (user *User) Save() *errors.ErrorRespond {

	// prepare statment for save new user to database
	stmt, err := mysql_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("internal server error %s ", err.Error()))
	}

	defer stmt.Close()

	// set user date created
	// use standard time zone
	// for Thailand need to +7 from UTC
	// now := time.Now().UTC()
	// user.DateCreated = now.Format("2006-01-02T15:04:05Z")
	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Avatar, user.Status, user.DateCreated)

	if err != nil {

		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s is already exist", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprint("internal server error %s ", err.Error()))
	}

	//user.DateCreated = date_utils.GetNow().Format("2006-01-02T15:04:05Z")
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprint("internal server error %s ", err.Error()))
	}

	user.Id = userId
	return nil
}
