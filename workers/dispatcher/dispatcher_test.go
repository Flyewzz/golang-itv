package dispatcher

import (
	"reflect"
	"testing"
	"time"

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
			name: "Dispatcher 1",
			d: NewDispatcher(2, 5, 5,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
		},
		{
			name: "Dispatcher 2",
			d: NewDispatcher(1, 5, 10,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
		},
		{
			name: "Dispatcher 3",
			d: NewDispatcher(10, 10, 10,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
		},
		{
			name: "Dispatcher 4",
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
		want *Result
	}{
		{
			name: "Dispatcher succ 1",
			d: NewDispatcher(2, 5, 5,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
			args: args{
				task: &Task{},
			},
			want: mocks.GetMockStandardSuccResult(),
		},
		{
			name: "Dispatcher succ 2",
			d: NewDispatcher(5, 1, 6,
				mocks.NewMockSuccessExecutor(), mocks.NewMockStoreController()),
			args: args{
				task: &Task{},
			},
			want: mocks.GetMockStandardSuccResult(),
		},
		{
			name: "Dispatcher fail 1",
			d: NewDispatcher(15, 1, 10,
				mocks.NewMockFailExecutor(), mocks.NewMockStoreController()),
			args: args{
				task: &Task{},
			},
			want: mocks.GetMockStandardFailResult(),
		},
		{
			name: "Dispatcher fail 2",
			d: NewDispatcher(5, 5, 5,
				mocks.NewMockFailExecutor(), mocks.NewMockStoreController()),
			args: args{
				task: &Task{},
			},
			want: mocks.GetMockStandardFailResult(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Dispatch()
			resCh := tt.d.AddNewTask(tt.args.task)
			timer := time.NewTimer(tt.d.timeout)
		chanCyc:
			for {
				select {
				case res := <-resCh:
					if !reflect.DeepEqual(res, tt.want) {
						t.Errorf("Dispatcher.AddNewTask() = %v, want %v", res, tt.want)
					}
					tt.d.Stop()
					break chanCyc
				case <-timer.C:
					t.Error("Time is up")
					break chanCyc
				}
			}

		})
	}
}
