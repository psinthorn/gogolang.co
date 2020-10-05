package contents

import models "github.com/psinthorn/gogolang.co/domain/errors"

var (
	contentDB = make(map[int64]*Content)
)

//
// Get content by ID
//
func (content *Content) Get() (*Content, *models.ErrorRespond) {
	return nil, nil
}

//
// Create new content
//

func (content *Content) Create() (*Content, *models.ErrorRespond) {
	return nil, nil
}
