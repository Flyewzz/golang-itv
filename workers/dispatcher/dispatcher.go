package dispatcher

import (
	"github.com/Flyewzz/golang-itv/interfaces"
	. "github.com/Flyewzz/golang-itv/models"
	"github.com/Flyewzz/golang-itv/workers"
	workerModels "github.com/Flyewzz/golang-itv/workers/models"
	"time"
)

type Dispatcher struct {
	workerPool      []*workers.Worker
	tasksQueue      chan workerModels.Job
	executor        interfaces.Executor
	storeController interfaces.Store
	// in seconds
	timeout time.Duration
}

func NewDispatcher(workersCount, maxTasks, tSeconds int, ex interfaces.Executor, sc interfaces.Store) *Dispatcher {
	tasksQueue := make(chan workerModels.Job, maxTasks)
	var workerList []*workers.Worker
	for i := 0; i < workersCount; i++ {
		w := workers.NewWorker(i, tasksQueue, ex, sc)
		workerList = append(workerList, w)
	}
	return &Dispatcher{
		workerPool: workerList,
		tasksQueue: tasksQueue,
		timeout:    time.Duration(tSeconds) * time.Second,
	}
}

func (d *Dispatcher) Dispatch() {
	for _, worker := range d.workerPool {
		worker.Start(d.timeout)
	}
}

func (d *Dispatcher) Stop() {
	close(d.tasksQueue)
}

func (d *Dispatcher) AddNewTask(task *Task) chan *Result {
	resCh := make(chan *Result)
	job := workerModels.NewJob(task, resCh)
	go func(job *workerModels.Job) {
		d.tasksQueue <- *job
	}(&job)
	return resCh
}
