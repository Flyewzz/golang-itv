package workers

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/Flyewzz/golang-itv/interfaces"
	"github.com/Flyewzz/golang-itv/models"
	workerModels "github.com/Flyewzz/golang-itv/workers/models"
)

type SuccessExecutor struct{}

func NewMockSuccessExecutor() *SuccessExecutor { return &SuccessExecutor{} }

func (ex *SuccessExecutor) Execute(client *http.Client, task *models.Task) (*models.Response, error) {
	return &models.Response{
		Status:        "200 OK",
		Headers:       "Header1: Header1",
		Body:          "body",
		ContentLength: 4,
	}, nil
}

type FailExecutor struct{}

func NewMockFailExecutor() *FailExecutor { return &FailExecutor{} }

func (ex *FailExecutor) Execute(client *http.Client, task *models.Task) (*models.Response, error) {
	return nil, errors.New("Some error... :(")
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
	tests := []struct {
		name string
		w    *Worker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Start()
		})
	}
}

func TestWorker_SendResult(t *testing.T) {
	type args struct {
		result *models.Result
		resCh  chan *models.Result
	}
	tests := []struct {
		name string
		w    *Worker
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.SendResult(tt.args.result, tt.args.resCh)
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
