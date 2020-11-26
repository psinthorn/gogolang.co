package users

import (
	"fmt"

	mysql_db "github.com/psinthorn/gogolang.co/datasources/mysql/users_db"
	"github.com/psinthorn/gogolang.co/domains/errors"
	date_utils "github.com/psinthorn/gogolang.co/utils/date"
	mysql_utils "github.com/psinthorn/gogolang.co/utils/mysql"
)

const (
	indexUniqueEmail    = "email_UNIQUE"
	queryInsertUser     = "INSERT INTO USERS(first_name, last_name, email, avatar, status, date_created) VALUES(?,?,?,?,?,?);"
	queryGetAllUsers    = "SELECT * FROM users ORDER BY id DESC"
	queryGetUserById    = "SELECT id, first_name, last_name, email, avatar, status, date_created FROM users WHERE id = ?"
	queryUpdateUserById = "UPDATE users SET first_name=?, last_name=?, email=?, avatar=?, status=? WHERE id=?;"
	queryDeleteUserById = "DELETE FROM users WHERE id=?"
	errorNoRows         = "no rows in result set"
)

var (
	usersDB = make(map[int64]*User)
)

// Save user to database
func (user *User) Save() *errors.ErrorRespond {

	// prepare statment for save new user to database
	stmt, err := mysql_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return mysql_utils.PareError(err)
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
		return mysql_utils.PareError(err)
	}

	//user.DateCreated = date_utils.GetNow().Format("2006-01-02T15:04:05Z")
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.PareError(err)
	}

	user.Id = userId
	return nil
}

// Get all users from database
func GetAll() ([]*User, *errors.ErrorRespond) {

	// ใช้ Prepare เพื่อตรวจสอบความถูกต้องของข้อมูลก่อนที่จะส่งไปทำการ  process ที่ server เพื่อลดการทำงาน process ที่ฝั่ง server
	var user *User
	var allUsers []*User
	stmt, err := mysql_db.Client.Prepare(queryGetAllUsers)
	if err != nil {
		return nil, mysql_utils.PareError(err)
	}
	defer stmt.Close()

	results, err := stmt.Query()

	if results.Next() {
		err := results.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Avatar)
		if err != nil {

			return nil, mysql_utils.PareError(err)
		}
		allUsers = append(allUsers, user)
	}
	fmt.Println(allUsers)
	return allUsers, nil
}

// Get user by id from database
func (user *User) Get() *errors.ErrorRespond {

	// ใช้ Prepare เพื่อตรวจสอบความถูกต้องของข้อมูลก่อนที่จะส่งไปทำการ  process ที่ server เพื่อลดการทำงาน process ที่ฝั่ง server
	stmt, err := mysql_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return mysql_utils.PareError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Avatar, &user.DateCreated, &user.Status); err != nil {
		return mysql_utils.PareError(err)
	}

	return nil
}

func (user *User) Update() *errors.ErrorRespond {
	stmt, err := mysql_db.Client.Prepare(queryUpdateUserById)
	if err != nil {
		return mysql_utils.PareError(err)
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Avatar, user.Status, user.Id)
	if err != nil {
		return mysql_utils.PareError(err)
	}
	// if err != nil {
	// 	return mysql_utils.PareError(err)
	// }
	return nil
}

func (user *User) Delete() *errors.ErrorRespond {
	stmt, err := mysql_db.Client.Prepare(queryDeleteUserById)
	if err != nil {
		return mysql_utils.PareError(err)
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql_utils.PareError(err)
	}
	return nil
}
