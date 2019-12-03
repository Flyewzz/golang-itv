package main

import (
	"log"
	"net/http"
)

type FetchTask struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

var FetchTaskChannel chan *FetchTask

// var fetchOne sync.Once

var FetchDone chan struct{}

// var doneOne sync.Once

// func GetFetchTaskChannel() *chan *FetchTask {
// 	fetchOne.Do(func() {
// 		fetchTaskChannel = make(chan *FetchTask)
// 	})
// 	return &fetchTaskChannel
// }

var ListOfTasks []*FetchTask

// var fetchTaskOne sync.Once

func FetchTaskWorker() {
	select {
	case task := <-FetchTaskChannel:
		_, err := SendRequest(&http.Client{}, task.Method, task.Url)
		if err != nil {
			log.Println(err)
		}
	case <-FetchDone:
		log.Println("FetchTask worker was correctly finished.")
		return
	}
}
