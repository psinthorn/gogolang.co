package contents

import "github.com/psinthorn/gogolang.co/domains/errors"

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

func (content *Content) Save() *errors.ErrorRespond {
	contentDB[content.Id] = content
	return nil
}
