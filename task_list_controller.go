package main

import (
	"errors"
	"sort"
	"sync"

	"github.com/spf13/viper"
)

type ListController struct {
	itemsPerPage int
	currentId    int
	tasks        map[int]*Task
}

var listController ListController
var listOnce sync.Once

func GetListController() *ListController {
	listOnce.Do(func() {
		listController = ListController{
			itemsPerPage: viper.GetInt("itemsPerPage"),
			currentId:    0,
			tasks:        make(map[int]*Task),
		}
	})
	return &listController
}

func (lc *ListController) AddNew(task *Task) int {
	lc.currentId++
	lc.tasks[lc.currentId] = task
	return lc.currentId
}

func (lc *ListController) GetAll() []Task {
	var tasks []Task
	var ids []int
	for id := range lc.tasks {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for _, id := range ids {
		tasks = append(tasks, *lc.tasks[id])
	}
	return tasks
}

func (lc *ListController) RemoveAll() {
	for id := range lc.tasks {
		delete(lc.tasks, id)
	}
}

func (lc *ListController) GetTasksByPage(page int) ([]Task, error) {
	tasks := lc.GetAll()
	itemsPerPage := lc.itemsPerPage
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

func (lc *ListController) GetById(id int) (*Task, error) {
	task, ok := lc.tasks[id]
	if !ok {
		return nil, errors.New("A task was not found")
	}
	return task, nil
}

func (lc *ListController) RemoveById(id int) error {
	_, ok := lc.tasks[id]
	if !ok {
		return errors.New("A task was not found")
	}
	delete(lc.tasks, id)
	return nil
}
