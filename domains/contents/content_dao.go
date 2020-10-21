package contents

import (
	"fmt"

	mysql_db "github.com/psinthorn/gogolang.co/datasources/mysql/users_db"
	"github.com/psinthorn/gogolang.co/domains/errors"
	date_utils "github.com/psinthorn/gogolang.co/utils"
)

const (
	queryInsertContent  = "INSERT INTO contents(title, sub_title, content, content_type, category, image, tags, author, status, date_created) VALUES(?,?,?,?,?,?,?,?,?,?);"
	queryGetContentById = "SELECT * FROM contents WHERE id = ?"
)

var (
	contentDB = make(map[int64]*Content)
)

//
// Get content by ID
//
func (content *Content) Get() *errors.ErrorRespond {

	// prepare statment
	stmt, err := mysql_db.Client.Prepare(queryGetContentById)
	// if error handle it
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("internal server error: %s", err.Error()))
	}

	result := stmt.QueryRow(content.Id)
	if err := result.Scan(&content.Id, &content.Title, &content.SubTitle, &content.Content, &content.ContentType, &content.Category, &content.Image, &content.Tags, &content.Author, &content.Status, &content.DateCreated); err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error on trying to get user id: %d errors: %s ", content.Id, err.Error()))
	}

	return nil
}

//
// Create new content
//

func (content *Content) Save() *errors.ErrorRespond {

	// Prepare statement
	stmt, err := mysql_db.Client.Prepare(queryInsertContent)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	// Close statment protect run out connection
	defer stmt.Close()

	content.DateCreated = date_utils.GetNowString()
	result, err := stmt.Exec(content.Title, content.SubTitle, content.Content, content.ContentType, content.Category, content.Image, content.Tags, content.Author, content.Status, content.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("internal server error %s", err.Error()))
	}

	contentId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("internal server error:  %s", err.Error()))
	}
	content.Id = contentId
	return nil
}
