package contents

import "github.com/psinthorn/gogolang.co/domain/errors"

var (
	contentDB = make(map[int64]*Content)
)

//
// Get content by ID
//
func (content *Content) Get() *errors.ErrorRespond {
	return nil
}

//
// Create new content
//

func (content *Content) Create() *errors.ErrorRespond {
	contentDB[content.Id] = content
	return nil
}
