package interfaces

import (
	. "github.com/Flyewzz/golang-itv/models"
)

type Dispatcher interface {
	Dispatch()
	Stop()
	AddNewTask(task *Task) chan *Result
}
