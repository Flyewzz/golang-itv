package models

import (
	"github.com/Flyewzz/golang-itv/models"
)

// A job for sending
type Job struct {
	Task     *models.Task
	ResultCh chan *models.Result
}

func NewJob(task *models.Task, resCh chan *models.Result) Job {
	return Job{
		Task:     task,
		ResultCh: resCh,
	}
}
