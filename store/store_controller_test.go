package store

import (
	"reflect"
	"sync"
	"testing"

	. "github.com/Flyewzz/golang-itv/models"
)

func newMockStoreController(itemsPerPage, startId int, requests map[int]*Request) *StoreController {
	return &StoreController{
		itemsPerPage: itemsPerPage,
		currentId:    startId,
		requests:     requests,
		mtx:          new(sync.Mutex),
	}
}

func getTestRequests() []Request {

	return []Request{
		{
			Task: &Task{
				Method: "GET",
				Url:    "http://google.ru",
			},
			Response: &Response{
				Status:        "200 OK",
				Headers:       "header1: header1",
				Body:          "<html>...</html>",
				ContentLength: 1200,
			},
		},
		{
			Task: &Task{
				Method: "POST",
				Url:    "http://yandex.ru",
			},
			Response: &Response{
				Status:        "403 Forbidden",
				Headers:       "header1: header1",
				Body:          "<html>...</html>",
				ContentLength: 1200,
			},
		},
		{
			Task: &Task{
				Method: "PUT",
				Url:    "http://rambler.ru",
			},
			Response: &Response{
				Status:        "404 Not Found",
				Headers:       "header1: header1",
				Body:          "<html>...</html>",
				ContentLength: 1200,
			},
		},
		{
			Task: &Task{
				Method: "DELETE",
				Url:    "http://yahoo.com",
			},
			Response: &Response{
				Status:        "418 I'm a teapot",
				Headers:       "header1: header1",
				Body:          "<html>...</html>",
				ContentLength: 1200,
			},
		},
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
	t.Skip()
	// type fields struct {
	// 	currentId int
	// 	requests  map[int]*Request
	// }
	// type args struct {
	// 	request *Request
	// }
	// tests := []struct {
	// 	name   string
	// 	fields fields
	// 	args   args
	// 	want   int
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		sc := newMockStoreController(
	// 			2,
	// 			tt.fields.currentId,
	// 			tt.fields.requests,
	// 		)
	// 		if got := sc.Add(tt.args.request); got != tt.want {
	// 			t.Errorf("StoreController.AddNew() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestStoreController_GetAll(t *testing.T) {
	type fields struct {
		currentId int
		requests  map[int]*Request
	}
	testRequests := getTestRequests()
	mapRequests := map[int]*Request{
		1: &testRequests[0],
		2: &testRequests[1],
		3: &testRequests[2],
		4: &testRequests[3],
	}
	tests := []struct {
		name   string
		fields fields
		want   []Request
	}{
		{
			name: "First adding",
			fields: fields{
				currentId: 4,
				requests:  mapRequests,
			},
			want: testRequests,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := newMockStoreController(
				2,
				tt.fields.currentId,
				tt.fields.requests,
			)
			if got := sc.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreController.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreController_GetById(t *testing.T) {
	testRequests := getTestRequests()
	mapRequests := map[int]*Request{
		1: &testRequests[0],
		2: &testRequests[1],
		3: &testRequests[2],
		4: &testRequests[3],
	}
	type fields struct {
		currentId int
		requests  map[int]*Request
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Request
		wantErr bool
	}{
		{
			name: "get 3",
			fields: fields{
				currentId: 4,
				requests:  mapRequests,
			},
			args: args{
				id: 3,
			},
			want:    mapRequests[3],
			wantErr: false,
		},

		{
			name: "get 2",
			fields: fields{
				currentId: 4,
				requests:  mapRequests,
			},
			args: args{
				id: 2,
			},
			want:    mapRequests[2],
			wantErr: false,
		},

		{
			name: "get 4",
			fields: fields{
				currentId: 4,
				requests:  mapRequests,
			},
			args: args{
				id: 4,
			},
			want:    mapRequests[4],
			wantErr: false,
		},

		{
			name: "get 0",
			fields: fields{
				currentId: 4,
				requests:  mapRequests,
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
				requests:  mapRequests,
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
			sc := newMockStoreController(
				2,
				tt.fields.currentId,
				tt.fields.requests,
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
	testRequests := getTestRequests()
	mapRequests := map[int]*Request{
		1: &testRequests[0],
		2: &testRequests[1],
		3: &testRequests[2],
		4: &testRequests[3],
	}
	type fields struct {
		currentId int
		requests  map[int]*Request
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
				requests:  mapRequests,
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
				requests:  mapRequests,
			},
			args: args{
				id: 15,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := newMockStoreController(
				2,
				tt.fields.currentId,
				tt.fields.requests,
			)
			if err := sc.RemoveById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("StoreController.RemoveById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, ok := tt.fields.requests[tt.args.id]; ok {
				t.Errorf("StoreController.RemoveById() expected 'ok' value: %v, but got %v", !ok, ok)
			}
		})
	}
}

func TestStoreController_GetByPage(t *testing.T) {
	// Items (requests) per page for testing
	itemsPerPage := 2
	type fields struct {
		itemsPerPage int
		currentId    int
		requests     map[int]*Request
	}
	type args struct {
		page int
	}
	testRequests := getTestRequests()

	mapRequests := map[int]*Request{
		1: &testRequests[0],
		2: &testRequests[1],
		3: &testRequests[2],
		4: &testRequests[3],
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Request
		wantErr bool
	}{
		{
			name: "1st page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    4,
				requests:     mapRequests,
			},
			args: args{
				page: 1,
			},
			want:    testRequests[:itemsPerPage],
			wantErr: false,
		},

		{
			name: "2nd page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    4,
				requests:     mapRequests,
			},
			args: args{
				page: 2,
			},
			want:    testRequests[itemsPerPage : itemsPerPage+itemsPerPage],
			wantErr: false,
		},
		{
			name: "3rd page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    4,
				requests:     mapRequests,
			},
			args: args{
				page: 3,
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "0 (zero) page",
			fields: fields{
				itemsPerPage: itemsPerPage,
				currentId:    4,
				requests:     mapRequests,
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
				currentId:    4,
				requests:     mapRequests,
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
			sc := newMockStoreController(
				tt.fields.itemsPerPage,
				tt.fields.currentId,
				tt.fields.requests,
			)
			got, err := sc.GetByPage(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreController.GetByPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreController.GetByPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreController_RemoveAll(t *testing.T) {
	type fields struct {
		itemsPerPage int
		currentId    int
		requests     map[int]*Request
	}

	testRequests := getTestRequests()

	mapRequests := map[int]*Request{
		1: &testRequests[0],
		2: &testRequests[1],
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
				requests:     mapRequests,
			},
		},
		{
			name: "Empty map",
			fields: fields{
				itemsPerPage: 2,
				currentId:    0,
				requests:     make(map[int]*Request),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := newMockStoreController(
				tt.fields.itemsPerPage,
				tt.fields.currentId,
				tt.fields.requests,
			)
			sc.RemoveAll()
			if len(sc.requests) != 0 {
				t.Errorf("Expected 0 requests, but got: %d\n", len(sc.requests))
			}
		})
	}
}
