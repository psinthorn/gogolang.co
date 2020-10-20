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

func GetAllContent(content contents.Content) (*contents.Content, *errors.ErrorRespond) {
	if err := content.Get(); err != nil {
		return nil, err
	}
	return &content, nil
}

func GetContent(id int64) (*contents.Content, *errors.ErrorRespond) {

	return nil, nil
}

func UpdateContent(content contents.Content) (*contents.Content, *errors.ErrorRespond) {
	return nil, nil
}

func DeleteContent(content contents.Content) (*contents.Content, *errors.ErrorRespond) {
	return nil, nil
}
