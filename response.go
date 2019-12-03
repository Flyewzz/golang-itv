package main

type Response struct {
	Id            int    `json:"id"`
	Status        string `json:"status"`
	Headers       string `json:"headers"`
	Body          string `json:"body"`
	ContentLength int64  `json:"content-length"`
}

func NewResponse(id int, status, headers, body string, contentLength int64) *Response {
	return &Response{
		Id:            id,
		Status:        status,
		Headers:       headers,
		Body:          body,
		ContentLength: contentLength,
	}
}
