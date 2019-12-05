package main

import (
	"fmt"
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
	for _, a := range mapTasks {
		fmt.Println(a.Url)
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
			if got := lc.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListController.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListController_GetById(t *testing.T) {
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
		want    *Task
		wantErr bool
	}{
		{
			name: "get 3",
			fields: fields{
				currentId: 4,
				tasks:     mapTasks,
			},
			args: args{
				id: 3,
			},
			want:    mapTasks[3],
			wantErr: false,
		},

		{
			name: "get 2",
			fields: fields{
				currentId: 4,
				tasks:     mapTasks,
			},
			args: args{
				id: 2,
			},
			want:    mapTasks[2],
			wantErr: false,
		},

		{
			name: "get 4",
			fields: fields{
				currentId: 4,
				tasks:     mapTasks,
			},
			args: args{
				id: 4,
			},
			want:    mapTasks[4],
			wantErr: false,
		},

		{
			name: "get 0",
			fields: fields{
				currentId: 4,
				tasks:     mapTasks,
			},
			args: args{
				id: 0,
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "get -5",
			fields: fields{
				currentId: 4,
				tasks:     mapTasks,
			},
			args: args{
				id: -5,
			},
			want:    nil,
			wantErr: true,
		},
	}
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

func TestListController_GetTasksByPage(t *testing.T) {
	// Items (tasks) per page for testing
	itemsPerPage := 2
	type fields struct {
		itemsPerPage int
		currentId    int
		tasks        map[int]*Task
	}
	type args struct {
		page int
	}
	testTasks := []Task{
		{
			Method: "GET",
			Url:    "Url1",
		},
		{
			Method: "POST",
			Url:    "Url2",
		},
		{
			Method: "PUT",
			Url:    "Url3",
		},
		{
			Method: "DELETE",
			Url:    "Url4",
		},
		{
			Method: "GET",
			Url:    "Url5",
		},
	}

	mapTasks := map[int]*Task{
		1: &testTasks[0],
		2: &testTasks[1],
		3: &testTasks[2],
		4: &testTasks[3],
		5: &testTasks[4],
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		{
			name: "1st page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    5,
				tasks:        mapTasks,
			},
			args: args{
				page: 1,
			},
			want:    testTasks[:itemsPerPage],
			wantErr: false,
		},

		{
			name: "2nd page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    5,
				tasks:        mapTasks,
			},
			args: args{
				page: 2,
			},
			want:    testTasks[itemsPerPage : itemsPerPage+itemsPerPage],
			wantErr: false,
		},
		{
			name: "3rd page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    5,
				tasks:        mapTasks,
			},
			args: args{
				page: 3,
			},
			want:    testTasks[itemsPerPage*2:],
			wantErr: false,
		},

		{
			name: "0 (zero) page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    5,
				tasks:        mapTasks,
			},
			args: args{
				page: 0,
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "-1 page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    5,
				tasks:        mapTasks,
			},
			args: args{
				page: -1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &ListController{
				itemsPerPage: tt.fields.itemsPerPage,
				currentId:    tt.fields.currentId,
				tasks:        tt.fields.tasks,
			}
			got, err := lc.GetTasksByPage(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListController.GetTasksByPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListController.GetTasksByPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
