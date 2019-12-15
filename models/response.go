package models

type Response struct {
	Status        string `json:"status"`
	Headers       string `json:"headers"`
	Body          string `json:"body"`
	ContentLength int64  `json:"content-length"`
}

func NewResponse(status, headers, body string, contentLength int64) *Response {
	return &Response{
		Status:        status,
		Headers:       headers,
		Body:          body,
		ContentLength: contentLength,
	}
}
