package errors

type ErrorRespond struct {
	//Id         string `json: "id"`
	Message    string `json: "message"`
	StatusCode int    `json: "status_code"`
	Error      string `json: "error"`
}
