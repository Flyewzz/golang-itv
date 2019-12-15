package store

import (
	"errors"
	. "github.com/Flyewzz/golang-itv/models"
	"sort"
	"sync"
)

type StoreController struct {
	itemsPerPage int
	currentId    int
	requests     map[int]*Request
	mtx          *sync.Mutex
}

func NewStoreController(itemsPerPage, startId int) *StoreController {
	return &StoreController{
		itemsPerPage: itemsPerPage,
		currentId:    startId,
		requests:     make(map[int]*Request),
		mtx:          new(sync.Mutex),
	}
}

func (sc *StoreController) GetAll() []Request {
	var Requests []Request
	var ids []int
	sc.mtx.Lock()
	for id := range sc.requests {
		ids = append(ids, id)
	}
	sc.mtx.Unlock()
	sort.Ints(ids)
	for _, id := range ids {
		Requests = append(Requests, *sc.requests[id])
	}
	return Requests
}

func (sc *StoreController) GetByPage(page int) ([]Request, error) {
	requests := sc.GetAll()
	itemsPerPage := sc.itemsPerPage
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage

	if start > (len(requests)-1) || start < 0 {
		return nil, errors.New("The incorrect page number.")
	}

	if stop > len(requests) {
		stop = len(requests)
	}

	return requests[start:stop], nil
}

func (sc *StoreController) GetById(id int) (*Request, error) {
	sc.mtx.Lock()
	Request, ok := sc.requests[id]
	sc.mtx.Unlock()
	if !ok {
		return nil, errors.New("A Request was not found")
	}
	return Request, nil
}

func (sc *StoreController) Add(Request *Request) int {
	sc.mtx.Lock()
	sc.currentId++
	sc.requests[sc.currentId] = Request
	sc.mtx.Unlock()
	return sc.currentId
}

func (sc *StoreController) RemoveById(id int) error {
	sc.mtx.Lock()
	_, ok := sc.requests[id]
	sc.mtx.Unlock()
	if !ok {
		return errors.New("A Request was not found")
	}
	sc.mtx.Lock()
	delete(sc.requests, id)
	sc.mtx.Unlock()
	return nil
}

func (sc *StoreController) RemoveAll() {
	sc.mtx.Lock()
	for id := range sc.requests {
		delete(sc.requests, id)
	}
	sc.mtx.Unlock()
}
