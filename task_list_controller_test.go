package main

import (
	"reflect"
	"testing"
)

func TestGetListController(t *testing.T) {
	tests := []struct {
		name string
		want *ListController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetListController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListController_AddNew(t *testing.T) {
	type fields struct {
		currentId int
		tasks     map[int]*Task
	}
	type args struct {
		task *Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &ListController{
				currentId: tt.fields.currentId,
				tasks:     tt.fields.tasks,
			}
			if got := lc.AddNew(tt.args.task); got != tt.want {
				t.Errorf("ListController.AddNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListController_GetAll(t *testing.T) {
	type fields struct {
		currentId int
		tasks     map[int]*Task
	}
	testTasks := []Task{
		Task{
			Method: "GET",
			Url:    "http://google.ru",
		},
		Task{
			Method: "POST",
			Url:    "http://yandex.ru",
		},
		Task{
			Method: "PUT",
			Url:    "http://rambler.ru",
		},
		Task{
			Method: "DELETE",
			Url:    "http://yahoo.com",
		},
	}
	mapTasks := map[int]*Task{
		1: &testTasks[0],
		2: &testTasks[1],
		3: &testTasks[2],
		4: &testTasks[3],
	}
	tests := []struct {
		name   string
		fields fields
		want   []Task
	}{
		{
			name: "First adding",
			fields: fields{
				currentId: 4,
				tasks:     mapTasks,
			},
			want: testTasks,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &ListController{
				currentId: tt.fields.currentId,
				tasks:     tt.fields.tasks,
			}
			if got := lc.GetAll(); !CompareSets(got, tt.want) {
				t.Errorf("ListController.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListController_GetById(t *testing.T) {

	// testTasks := []Task{
	// 	Task{
	// 		Method: "GET",
	// 		Url:    "http://google.ru",
	// 	},
	// 	Task{
	// 		Method: "POST",
	// 		Url:    "http://yandex.ru",
	// 	},
	// 	Task{
	// 		Method: "PUT",
	// 		Url:    "http://rambler.ru",
	// 	},
	// 	Task{
	// 		Method: "DELETE",
	// 		Url:    "http://yahoo.com",
	// 	},
	// }
	// mapTasks := map[int]*Task{
	// 	1: &testTasks[0],
	// 	2: &testTasks[1],
	// 	3: &testTasks[2],
	// 	4: &testTasks[3],
	// }
	type fields struct {
		currentId int
		tasks     map[int]*Task
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Task
		wantErr bool
	}{}
	// errors.New("A task was not found")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &ListController{
				currentId: tt.fields.currentId,
				tasks:     tt.fields.tasks,
			}
			got, err := lc.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListController.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListController.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListController_RemoveById(t *testing.T) {
	testTasks := []Task{
		Task{
			Method: "GET",
			Url:    "http://google.ru",
		},
		Task{
			Method: "POST",
			Url:    "http://yandex.ru",
		},
		Task{
			Method: "PUT",
			Url:    "http://rambler.ru",
		},
		Task{
			Method: "DELETE",
			Url:    "http://yahoo.com",
		},
	}
	mapTasks := map[int]*Task{
		1: &testTasks[0],
		2: &testTasks[1],
		3: &testTasks[2],
		4: &testTasks[3],
	}
	type fields struct {
		currentId int
		tasks     map[int]*Task
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "First",
			fields: fields{
				currentId: 4,
				tasks:     mapTasks,
			},
			args: args{
				id: 2,
			},
			wantErr: false,
		},
		{
			name: "Second",
			fields: fields{
				currentId: 4,
				tasks:     mapTasks,
			},
			args: args{
				id: 15,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &ListController{
				currentId: tt.fields.currentId,
				tasks:     tt.fields.tasks,
			}
			if err := lc.RemoveById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ListController.RemoveById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, ok := tt.fields.tasks[tt.args.id]; ok {
				t.Errorf("ListController.RemoveById() expected 'ok' value: %v, but got %v", !ok, ok)
			}
		})
	}
}
