package models

type Task struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

func NewTask(method, url string) *Task {
	return &Task{
		Method: method,
		Url:    url,
	}
}
