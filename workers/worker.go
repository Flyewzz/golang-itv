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

func (w *Worker) Start(timeout time.Duration) {
	go func(w *Worker) {
		for job := range w.jobChannel {
			resp, err := w.exeсutor.Execute(&http.Client{
				Timeout: timeout,
			}, job.Task)
			result := models.NewResult(resp, err)
			type data struct {
				w      *Worker
				result *models.Result
				job    *workerModels.Job
				err    error
			}
			sending := &data{
				w:      w,
				result: result,
				job:    &job,
				err:    err,
			}
			go func(data *data) {
				// Before sending we should save request to the storage
				// (if there wasn't an error)
				if data.err == nil {
					// Save only corrected results
					data.w.SaveResult(data.job, data.result)
				}
				data.w.SendResult(data.result, data.job.ResultCh)
				// Make a finished log only after done the task
				data.w.logger.MakeFinishedLog(w.ID)
			}(sending)
		}
	}(w)
}

// Send a done job to its result channel
func (w *Worker) SendResult(result *models.Result, resCh chan<- *models.Result) {
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
