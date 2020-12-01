package contents

import (
	mysql_db "github.com/psinthorn/gogolang.co/datasources/mysql/users_db"
	"github.com/psinthorn/gogolang.co/domains/errors"
	date_utils "github.com/psinthorn/gogolang.co/utils/date"
	mysql_utils "github.com/psinthorn/gogolang.co/utils/mysql"
)

const (
	queryInsertContent     = "INSERT INTO contents(title, sub_title, content, content_type, category, image, tags, author, status, date_created) VALUES(?,?,?,?,?,?,?,?,?,?);"
	queryGetContentById    = "SELECT * FROM contents WHERE id = ?"
	queryDeleteContentById = "DELETE FROM contents where id = ?"
	queryGetAllContents    = "SELECT * FROM contents ORDER BY id DESC"
)

var (
	contentDB = make(map[int64]*Content)
)

//
// Get All content by ID
//

func GetAll() ([]Content, *errors.ErrorRespond) {

	//prepar statments
	stmt, err := mysql_db.Client.Prepare(queryGetAllContents)
	// if error handle it
	if err != nil {
		panic(err.Error())
	}
	// Close statment protect run out connection
	defer stmt.Close()

	results, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}

	content := Content{}
	allContents := []Content{}

	for results.Next() {
		err := results.Scan(&content.Id, &content.Title, &content.SubTitle, &content.Content, &content.ContentType, &content.Category, &content.Image, &content.Tags, &content.Author, &content.Status, &content.DateCreated)
		if err != nil {
			panic(err.Error())
		}
		allContents = append(allContents, content)
	}

	return allContents, nil

}

//
// Get content by ID
//
func (content *Content) Get() *errors.ErrorRespond {

	// prepare statment
	stmt, err := mysql_db.Client.Prepare(queryGetContentById)
	// if error handle it
	if err != nil {
		mysql_utils.PareError(err)
	}

	// Close statment protect run out connection
	defer stmt.Close()

	result := stmt.QueryRow(content.Id)
	if err := result.Scan(&content.Id, &content.Title, &content.SubTitle, &content.Content, &content.ContentType, &content.Category, &content.Image, &content.Tags, &content.Author, &content.Status, &content.DateCreated); err != nil {
		mysql_utils.PareError(err)
	}

	return nil
}

//
// Create new content
//

func (content *Content) Save() *errors.ErrorRespond {
	stmt, err := mysql_db.Client.Prepare(queryInsertContent)
	if err != nil {
		mysql_utils.PareError(err)
	}

	// Close statment protect run out connection
	defer stmt.Close()

	content.DateCreated = date_utils.GetNowString()
	result, err := stmt.Exec(content.Title, content.SubTitle, content.Content, content.ContentType, content.Category, content.Image, content.Tags, content.Author, content.Status, content.DateCreated)
	if err != nil {
		mysql_utils.PareError(err)
	}

	contentId, err := result.LastInsertId()
	if err != nil {
		mysql_utils.PareError(err)
	}
	content.Id = contentId
	return nil
}

func (content *Content) Delete() *errors.ErrorRespond {
	stmt, err := mysql_db.Client.Prepare(queryDeleteContentById)
	if err != nil {
		mysql_utils.PareError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(content.Id)
	if err != nil {
		return mysql_utils.PareError(err)
	}

	return nil

}
