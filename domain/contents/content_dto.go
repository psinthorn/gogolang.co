package contents

import (
	"strings"

	"github.com/psinthorn/gogolang.co/domain/errors"
)

type Content struct {
	Id          int64  `json: "id"`
	Title       string `json: "title"`
	Content     string `json: "content"`
	ContentType string `json: "content_type"`
	Category    string `json: "category"`
	Author      string `json: "content"`
	Status      bool   `json: "status"`
	DateCreated string `json: "date_created"`
}

//
// Validate content input
//

func (content *Content) Validate() *errors.ErrorRespond {
	content.Title = strings.TrimSpace(content.Title)
	if content.Title == "" {
		return errors.NewContentAlertNotice("Title can not be empty!")
	}
	return nil
}
