package models

import "log"

type WorkerLogger struct{}

func (wl *WorkerLogger) MakeExecuteLog(id int) {
	log.Printf("Worker #%d was completed a task\n", id)
}

func (wl *WorkerLogger) MakeFinishedLog(id int) {
	log.Printf("Worker #%d was stopped\n", id)
}
