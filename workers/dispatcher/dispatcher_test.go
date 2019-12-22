package dispatcher

import (
	"reflect"
	"testing"

	"github.com/Flyewzz/golang-itv/features"
	"github.com/Flyewzz/golang-itv/mocks"
	. "github.com/Flyewzz/golang-itv/models"
)

func TestDispatcher_Dispatch(t *testing.T) {
	tests := []struct {
		name string
		d    *Dispatcher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Dispatch()
		})
	}
}

func TestDispatcher_Stop(t *testing.T) {
	tests := []struct {
		name string
		d    *Dispatcher
	}{
		{
			name: "Dispacher 1",
			d: NewDispatcher(2, 5, 5,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
		},
		{
			name: "Dispacher 2",
			d: NewDispatcher(1, 5, 10,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
		},
		{
			name: "Dispacher 3",
			d: NewDispatcher(10, 10, 10,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
		},
		{
			name: "Dispacher 4",
			d: NewDispatcher(5, 1, 1,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Stop()
			closed := features.IsClosed(tt.d.tasksQueue)
			if !closed {
				t.Errorf("Expected closed: %v, but got: %v",
					closed, !closed)
			}
		})
	}
}

func TestDispatcher_AddNewTask(t *testing.T) {
	type args struct {
		task *Task
	}
	tests := []struct {
		name string
		d    *Dispatcher
		args args
		want chan *Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddNewTask(tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.AddNewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
