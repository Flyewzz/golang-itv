package models

// A worker will return the result with an error to HandlerData
type Result struct {
	Response *Response
	Error    error
}

func NewResult(resp *Response, err error) *Result {
	return &Result{
		Response: resp,
		Error:    err,
	}
}
