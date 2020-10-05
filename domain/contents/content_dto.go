package contents

import (
	"strings"

	"github.com/psinthorn/gogolang.co/domain/errors"
)

type Content struct {
	Id          int64  `json: "id"`
	Title       string `json: "title"`
	Content     string `json: "content"`
	Author      string `json: "content"`
	DateCreated string `json: "date_created"`
	Status      bool   `json: "status"`
}

//
// Validate content input
//

func (content *Content) Validate() *errors.ErrorRespond {
	content.Title = strings.TrimSpace(content.Title)
	if content.Title == "" {
		err := errors.NewContentAlertNotice("Title Can not be blank!")
		return err
	}
	return nil
}
