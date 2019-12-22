package models

type Request struct {
	Task     *Task     `json:"task"`
	Response *Response `json:"response"`
}

func NewRequest(t *Task, r *Response) *Request {
	return &Request{
		Task:     t,
		Response: r,
	}
}
