package main

import (
	"errors"
	"sync"
)

type ListController struct {
	currentId int
	tasks     map[int]*Task
}

var listController ListController
var listOnce sync.Once

func GetListController() *ListController {
	listOnce.Do(func() {
		listController = ListController{
			currentId: 0,
			tasks:     make(map[int]*Task),
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
	for _, task := range lc.tasks {
		tasks = append(tasks, *task)
	}
	return tasks
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
