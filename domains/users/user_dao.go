package users

import (
	"fmt"

	mysql_db "github.com/psinthorn/gogolang.co/datasources/mysql/users_db"
	"github.com/psinthorn/gogolang.co/domains/errors"
	mysql_utils "github.com/psinthorn/gogolang.co/utils/mysql"
)

const (
	indexUniqueEmail      = "email_UNIQUE"
	queryInsertUser       = "INSERT INTO USERS(first_name, last_name, email, password, avatar, status, date_created) VALUES(?,?,?,?,?,?,?);"
	queryGetAllUsers      = "SELECT id, first_name, last_name, email, avatar, status, date_created FROM users ORDER BY id DESC"
	queryGetUserById      = "SELECT id, first_name, last_name, email, avatar, status, date_created FROM users WHERE id = ?"
	queryUpdateUserById   = "UPDATE users SET first_name=?, last_name=?, email=?, avatar=?, status=? WHERE id=?;"
	queryDeleteUserById   = "DELETE FROM users WHERE id=?"
	errorNoRows           = "no rows in result set"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, avatar, date_created, status FROM users WHERE status=?"
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

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Avatar, user.Status, user.DateCreated)

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
func (user *User) GetAll() ([]User, *errors.ErrorRespond) {
	// ใช้ Prepare เพื่อตรวจสอบความถูกต้องของข้อมูลก่อนที่จะส่งไปทำการ  process ที่ server เพื่อลดการทำงาน process ที่ฝั่ง server
	stmt, err := mysql_db.Client.Prepare(queryGetAllUsers)
	if err != nil {
		return nil, mysql_utils.PareError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, mysql_utils.PareError(err)
	}
	defer rows.Close()
	// fmt.Println(results)

	results := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Avatar, &user.Status, &user.DateCreated)
		if err != nil {
			return nil, mysql_utils.PareError(err)
		}
		results = append(results, user)
	}

	// if results == 0 then return err not found
	if len(results) == 0 {
		return nil, errors.NewNotFoundError("No user found")
	}
	// Retrun results to request
	return results, nil
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
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Avatar, &user.Status, &user.DateCreated); err != nil {
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

func (user *User) FindUserByStatus(status string) ([]User, *errors.ErrorRespond) {
	stmt, err := mysql_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, mysql_utils.PareError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.PareError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Avatar, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.PareError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
