package models

import "log"

type WorkerLogger struct{}

func (wl *WorkerLogger) MakeExecuteLog(id int) {
	log.Printf("Worker #%d was completed a task\n", id)
}
