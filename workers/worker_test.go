package workers

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Flyewzz/golang-itv/interfaces"
	"github.com/Flyewzz/golang-itv/models"
	workerModels "github.com/Flyewzz/golang-itv/workers/models"
)

var (
	FAIL_EXECUTOR_ERROR error = errors.New("Some error... :(")
)

type SuccessExecutor struct{}

func NewMockSuccessExecutor() *SuccessExecutor { return &SuccessExecutor{} }

func getStandardSuccResponse() *models.Response {
	return &models.Response{
		Status:        "200 OK",
		Headers:       "Header1: Header1",
		Body:          "body",
		ContentLength: 4,
	}
}

func (ex *SuccessExecutor) Execute(client *http.Client, task *models.Task) (*models.Response, error) {
	return getStandardSuccResponse(), nil
}

type FailExecutor struct{}

func NewMockFailExecutor() *FailExecutor { return &FailExecutor{} }

func (ex *FailExecutor) Execute(client *http.Client, task *models.Task) (*models.Response, error) {
	return nil, FAIL_EXECUTOR_ERROR
}

type MockStoreController struct {
	last *models.Request
}

func NewMockStoreController() *MockStoreController {
	return &MockStoreController{}
}

func (st *MockStoreController) Add(request *models.Request) int { st.last = request; return 0 }
func (st *MockStoreController) GetAll() []models.Request        { return []models.Request{} }
func (st *MockStoreController) GetByPage(page int) ([]models.Request, error) {
	return []models.Request{}, nil
}
func (st *MockStoreController) GetById(id int) (*models.Request, error) {
	if id == 0 {
		return st.last, nil
	}
	return nil, errors.New("Incorrect index")
}
func (st *MockStoreController) RemoveById(id int) error { return nil }
func (st *MockStoreController) RemoveAll()              {}

func TestNewWorker(t *testing.T) {
	type args struct {
		id int
		tc chan workerModels.Job
		ex interfaces.Executor
		sc interfaces.Store
	}
	var ex interfaces.Executor = NewMockSuccessExecutor()
	var sc interfaces.Store = NewMockStoreController()
	jobCh := make(chan workerModels.Job)

	tests := []struct {
		name string
		args args
		want *Worker
	}{
		{
			name: "Worker1",
			args: args{
				id: 0,
				tc: jobCh,
				ex: ex,
				sc: sc,
			},
			want: &Worker{
				ID:              0,
				jobChannel:      jobCh,
				exeсutor:        ex,
				storeController: sc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWorker(tt.args.id, tt.args.tc, tt.args.ex, tt.args.sc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWorker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorker_Start(t *testing.T) {
	succEx := NewMockSuccessExecutor()
	failEx := NewMockFailExecutor()
	storeController := NewMockStoreController()
	jobsChannels := []chan workerModels.Job{
		make(chan workerModels.Job, 10),
		make(chan workerModels.Job, 10),
		make(chan workerModels.Job, 10),
	}
	resultChannels := []chan *models.Result{
		make(chan *models.Result),
		make(chan *models.Result),
		make(chan *models.Result),
	}
	tests := []struct {
		name  string
		w     *Worker
		resCh chan *models.Result
		job   workerModels.Job
		want  *models.Result
	}{
		{
			name:  "OK worker1",
			w:     NewWorker(0, jobsChannels[0], succEx, storeController),
			resCh: resultChannels[0],
			job:   workerModels.NewJob(&models.Task{}, resultChannels[0]),
			want: &models.Result{
				Response: getStandardSuccResponse(),
				Error:    nil,
			},
		},
		{
			name:  "OK worker2",
			w:     NewWorker(1, jobsChannels[1], succEx, storeController),
			resCh: resultChannels[1],
			job:   workerModels.NewJob(&models.Task{}, resultChannels[1]),
			want: &models.Result{
				Response: getStandardSuccResponse(),
				Error:    nil,
			},
		},
		{
			name:  "FAIL worker3",
			w:     NewWorker(2, jobsChannels[2], failEx, storeController),
			resCh: resultChannels[2],
			job:   workerModels.NewJob(&models.Task{}, resultChannels[2]),
			want: &models.Result{
				Response: nil,
				Error:    FAIL_EXECUTOR_ERROR,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeout := 5 * time.Second
			timer := time.NewTimer(timeout)
			tt.w.Start(timeout)
			go func(job *workerModels.Job) { tt.w.jobChannel <- *job }(&tt.job)
		chanCyc:
			for {
				select {
				case res := <-tt.resCh:
					if !reflect.DeepEqual(res, tt.want) {
						t.Error("Results are different")
					}
					break chanCyc
				case <-timer.C:
					t.Error("Time is up")
					break chanCyc
				}
			}
		})
	}
}

func TestWorker_SendResult(t *testing.T) {
	jobCh := make(chan workerModels.Job, 1)
	succEx := NewMockSuccessExecutor()
	failEx := NewMockFailExecutor()
	storeController := NewMockStoreController()
	type args struct {
		result *models.Result
		resCh  chan *models.Result
	}
	tests := []struct {
		name    string
		w       *Worker
		args    args
		wantErr bool
	}{
		{
			name: "get google",
			w:    NewWorker(0, jobCh, succEx, storeController),
			args: args{
				result: &models.Result{
					Response: &models.Response{
						Status:        "200 OK",
						Headers:       "Header1: header1",
						Body:          "body",
						ContentLength: 4,
					},
					Error: nil,
				},
				resCh: make(chan *models.Result),
			},
			wantErr: false,
		},
		{
			name: "post yandex",
			w:    NewWorker(1, jobCh, succEx, storeController),
			args: args{
				result: &models.Result{
					Response: &models.Response{
						Status:        "201 Created",
						Headers:       "Header1: header1\nHeader2: header2",
						Body:          "body and tody",
						ContentLength: 13,
					},
					Error: nil,
				},
				resCh: make(chan *models.Result),
			},
			wantErr: false,
		},
		{
			name: "delete rambler.ru",
			w:    NewWorker(2, jobCh, failEx, storeController),
			args: args{
				result: &models.Result{
					Response: nil,
					Error:    errors.New("Some error was occured"),
				},
				resCh: make(chan *models.Result),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go tt.w.SendResult(tt.args.result, tt.args.resCh)
			timeout := time.NewTimer(15 * time.Second)
		chanCyc:
			for {
				select {
				case res := <-tt.args.resCh:
					isErr := (res.Error != nil)
					if isErr != tt.wantErr {
						t.Errorf("Expected error: %v, but get: %v", res.Error, tt.wantErr)
						break chanCyc
					}
					if !reflect.DeepEqual(res, tt.args.result) {
						t.Errorf("Expected result: \n"+
							"	Response: \n"+
							"		status: %s\n"+
							"		headers: %s\n"+
							"		body: %s\n"+
							"		content-length: %d\n"+
							"	Error: %v \n\n"+
							"Actual result: \n"+
							"	Response: \n"+
							"		status: %s\n"+
							"		headers: %s\n"+
							"		body: %s\n"+
							"		content-length: %d\n"+
							"	Error: %v \n\n",
							tt.args.result.Response.Status,
							tt.args.result.Response.Headers,
							tt.args.result.Response.Body,
							tt.args.result.Response.ContentLength,
							tt.args.result.Error,
							res.Response.Status,
							res.Response.Headers,
							res.Response.Body,
							res.Response.ContentLength,
							res.Error,
						)
					}
					break chanCyc
				case <-timeout.C:
					t.Error("Time is up")
					break chanCyc
				}
			}
		})
	}
}

func TestWorker_SaveResult(t *testing.T) {
	jobCh := make(chan workerModels.Job)
	type args struct {
		job    *workerModels.Job
		result *models.Result
	}
	tests := []struct {
		name string
		w    *Worker
		args args
		want *models.Request
	}{
		{
			name: "get google",
			w: NewWorker(0, jobCh, NewMockSuccessExecutor(),
				NewMockStoreController()),
			args: args{
				job: &workerModels.Job{
					Task: &models.Task{
						Method: "GET",
						Url:    "https://google.ru",
					},
				},
				result: &models.Result{
					Response: models.NewResponse("200 OK", "Header1: header1",
						"body", 4),
					Error: nil,
				},
			},
			want: &models.Request{
				Task: &models.Task{
					Method: "GET",
					Url:    "https://google.ru",
				},
				Response: models.NewResponse("200 OK", "Header1: header1",
					"body", 4),
			},
		},
		{
			name: "post yandex",
			w: NewWorker(0, jobCh, NewMockSuccessExecutor(),
				NewMockStoreController()),
			args: args{
				job: &workerModels.Job{
					Task: &models.Task{
						Method: "POST",
						Url:    "https://yandex.ru",
					},
				},
				result: &models.Result{
					Response: models.NewResponse("200 OK", "Header1: header1",
						"body", 4),
					Error: nil,
				},
			},
			want: &models.Request{
				Task: &models.Task{
					Method: "POST",
					Url:    "https://yandex.ru",
				},
				Response: models.NewResponse("200 OK", "Header1: header1",
					"body", 4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.SaveResult(tt.args.job, tt.args.result)
			stored, err := tt.w.storeController.GetById(0)
			if err != nil {
				t.Error("Unexpected error with getById(0)")
			}
			if !reflect.DeepEqual(tt.want, stored) {
				t.Errorf("Test '%s' was failed. "+
					"stored task: method: %s, url: %s; "+
					"actual task: method: %s, url: %s\n"+
					"stored response: status: %s, headers: %s, body: %s, Content-Length: %d\n"+
					"actual response: status: %s, headers: %s, body: %s, Content-Length: %d\n", tt.name,
					stored.Task.Method, stored.Task.Url, tt.want.Task.Method, tt.want.Task.Url,
					stored.Response.Status, stored.Response.Headers, stored.Response.Body, stored.Response.ContentLength,
					tt.want.Response.Status, tt.want.Response.Headers, tt.want.Response.Body, tt.want.Response.ContentLength)
			}
		})
	}
}
