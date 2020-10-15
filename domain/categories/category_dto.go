package categories

type Category struct {
	Id          int64  `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Status      string `json: "status"`
	Author      string `json: "author"`
	DateCreated string `json: "description"`
}
