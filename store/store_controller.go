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
	tasks        map[int]*Task
	mtx          *sync.Mutex
}

func NewStoreController(itemsPerPage, startId int) *StoreController {
	return &StoreController{
		itemsPerPage: itemsPerPage,
		currentId:    startId,
		tasks:        make(map[int]*Task),
		mtx:          new(sync.Mutex),
	}
}

func (sc *StoreController) GetAll() []Task {
	var tasks []Task
	var ids []int
	sc.mtx.Lock()
	for id := range sc.tasks {
		ids = append(ids, id)
	}
	sc.mtx.Unlock()
	sort.Ints(ids)
	for _, id := range ids {
		tasks = append(tasks, *sc.tasks[id])
	}
	return tasks
}

func (sc *StoreController) GetTasksByPage(page int) ([]Task, error) {
	tasks := sc.GetAll()
	itemsPerPage := sc.itemsPerPage
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage

	if start > len(tasks) || start < 0 {
		return nil, errors.New("The incorrect page number.")
	}

	if stop > len(tasks) {
		stop = len(tasks)
	}

	return tasks[start:stop], nil
}

func (sc *StoreController) GetById(id int) (*Task, error) {
	sc.mtx.Lock()
	task, ok := sc.tasks[id]
	sc.mtx.Unlock()
	if !ok {
		return nil, errors.New("A task was not found")
	}
	return task, nil
}

func (sc *StoreController) Add(task *Task) int {
	sc.mtx.Lock()
	sc.currentId++
	sc.tasks[sc.currentId] = task
	sc.mtx.Unlock()
	return sc.currentId
}

func (sc *StoreController) RemoveById(id int) error {
	sc.mtx.Lock()
	_, ok := sc.tasks[id]
	sc.mtx.Unlock()
	if !ok {
		return errors.New("A task was not found")
	}
	sc.mtx.Lock()
	delete(sc.tasks, id)
	sc.mtx.Unlock()
	return nil
}

func (sc *StoreController) RemoveAll() {
	sc.mtx.Lock()
	for id := range sc.tasks {
		delete(sc.tasks, id)
	}
	sc.mtx.Unlock()
}
