package workers

import (
	"log"
	"net/http"
	"time"

	"github.com/Flyewzz/golang-itv/interfaces"
	"github.com/Flyewzz/golang-itv/models"
	workerModels "github.com/Flyewzz/golang-itv/workers/models"
)

type Worker struct {
	ID              int
	jobChannel      chan workerModels.Job
	exeсutor        interfaces.Executor
	storeController interfaces.Store
	logger          workerModels.WorkerLogger
}

func NewWorker(id int, tc chan workerModels.Job, ex interfaces.Executor, sc interfaces.Store) *Worker {
	return &Worker{
		ID:              id,
		jobChannel:      tc,
		exeсutor:        ex,
		storeController: sc,
	}
}

func (w *Worker) Start() {
	go func() {
		timeout := 5 * time.Second
		for job := range w.jobChannel {
			resp, err := w.exeсutor.Execute(&http.Client{
				Timeout: timeout,
			}, job.Task)
			result := models.NewResult(resp, err)
			go func(w *Worker, result *models.Result, job *workerModels.Job) {
				// Before sending we should save request to the storage
				w.SaveResult(job, result)
				w.SendResult(result, job.ResultCh)
			}(w, result, &job)
		}
	}()
}

// Send a done job to its result channel
func (w *Worker) SendResult(result *models.Result, resCh chan *models.Result) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error to send a completed job: %v\n", r)
		}
	}()
	resCh <- result
	w.logger.MakeExecuteLog(w.ID)
}

// Save request to the storage
func (w *Worker) SaveResult(job *workerModels.Job, result *models.Result) {
	request := models.NewRequest(job.Task, result.Response)
	w.storeController.Add(request)
}
