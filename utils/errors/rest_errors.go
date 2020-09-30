package utils

type restError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}
