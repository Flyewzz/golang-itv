package store

import (
	"reflect"
	"sync"
	"testing"

	. "github.com/Flyewzz/golang-itv/models"
)

func NewMockStoreController(itemsPerPage, startId int, tasks map[int]*Task) *StoreController {
	return &StoreController{
		itemsPerPage: itemsPerPage,
		currentId:    startId,
		tasks:        tasks,
		mtx:          new(sync.Mutex),
	}
}

func TestGetStoreController(t *testing.T) {
	t.Skip()
	// tests := []struct {
	// 	name string
	// 	want *StoreController
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := NewStoreController(); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("NewStoreController() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestStoreController_Add(t *testing.T) {
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
			sc := NewMockStoreController(
				2,
				tt.fields.currentId,
				tt.fields.tasks,
			)
			if got := sc.Add(tt.args.task); got != tt.want {
				t.Errorf("StoreController.AddNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreController_GetAll(t *testing.T) {
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
			sc := NewMockStoreController(
				2,
				tt.fields.currentId,
				tt.fields.tasks,
			)
			if got := sc.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreController.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreController_GetById(t *testing.T) {
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
			sc := NewMockStoreController(
				2,
				tt.fields.currentId,
				tt.fields.tasks,
			)
			got, err := sc.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreController.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreController.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreController_RemoveById(t *testing.T) {
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
			sc := NewMockStoreController(
				2,
				tt.fields.currentId,
				tt.fields.tasks,
			)
			if err := sc.RemoveById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("StoreController.RemoveById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, ok := tt.fields.tasks[tt.args.id]; ok {
				t.Errorf("StoreController.RemoveById() expected 'ok' value: %v, but got %v", !ok, ok)
			}
		})
	}
}

func TestStoreController_GetTasksByPage(t *testing.T) {
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
			sc := NewMockStoreController(
				tt.fields.itemsPerPage,
				tt.fields.currentId,
				tt.fields.tasks,
			)
			got, err := sc.GetTasksByPage(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreController.GetTasksByPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreController.GetTasksByPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreController_RemoveAll(t *testing.T) {
	type fields struct {
		itemsPerPage int
		currentId    int
		tasks        map[int]*Task
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
	}

	mapTasks := map[int]*Task{
		1: &testTasks[0],
		2: &testTasks[1],
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Non-empty map",
			fields: fields{
				itemsPerPage: 2,
				currentId:    2,
				tasks:        mapTasks,
			},
		},
		{
			name: "Empty map",
			fields: fields{
				itemsPerPage: 2,
				currentId:    0,
				tasks:        make(map[int]*Task),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := NewMockStoreController(
				tt.fields.itemsPerPage,
				tt.fields.currentId,
				tt.fields.tasks,
			)
			sc.RemoveAll()
			if len(sc.tasks) != 0 {
				t.Errorf("Expected 0 tasks, but got: %d\n", len(sc.tasks))
			}
		})
	}
}
