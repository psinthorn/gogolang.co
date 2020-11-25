package services

import (
	"github.com/psinthorn/gogolang.co/domains/contents"
	"github.com/psinthorn/gogolang.co/domains/errors"
)

//
// Service connect to data access object (dao) work with data method on each request operation
//

func CreateContent(content contents.Content) (*contents.Content, *errors.ErrorRespond) {
	if err := content.Save(); err != nil {
		return nil, err
	}

	return &content, nil

}

func GetAllContent() (*contents.Content, *errors.ErrorRespond) {
	results := &contents.Content{}
	if err := results.GetAll(); err != nil {
		return nil, err
	}
	return results, nil
}

func GetContent(id int64) (*contents.Content, *errors.ErrorRespond) {
	result := &contents.Content{Id: id}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateContent(content contents.Content) (*contents.Content, *errors.ErrorRespond) {
	return nil, nil
}

func DeleteContent(id int64) *errors.ErrorRespond {
	content := &contents.Content{Id: id}
	return content.Delete()
}
