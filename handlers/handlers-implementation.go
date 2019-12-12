package handlers

import (
	"encoding/json"
	"fmt"
	. "github.com/Flyewzz/golang-itv/features"
	. "github.com/Flyewzz/golang-itv/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (uh *UserHandler) RequestHandler(w http.ResponseWriter, r *http.Request) {
	method, url := r.URL.Query().Get("method"), r.URL.Query().Get("url")
	if !CheckMethodValid(method) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	respId := uh.StoreController.Add(&Task{
		Method: method,
		Url:    url,
	})
	timeout := 5 * time.Second
	resp, err := uh.Executor.Execute(&http.Client{
		Timeout: timeout,
	}, method, url, respId)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(data)
}

func (uh *UserHandler) AllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := uh.StoreController.GetAll()
	data, _ := json.Marshal(tasks)
	w.Write(data)
}

func (uh *UserHandler) AllTasksRemoveHandler(w http.ResponseWriter, r *http.Request) {
	uh.StoreController.RemoveAll()
	w.Write([]byte("All tasks was successfully deleted."))
}

func (uh *UserHandler) PageTasksHandler(w http.ResponseWriter, r *http.Request) {
	strNum := mux.Vars(r)["number"]
	num, err := strconv.Atoi(strNum)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	tasks, err := uh.StoreController.GetTasksByPage(num)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(tasks)
	w.Write(data)
}

func (uh *UserHandler) TaskIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	task, err := uh.StoreController.GetById(numId)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(task)
	w.Write(data)
}

func (uh *UserHandler) RemoveTaskIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = uh.StoreController.RemoveById(numId)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Write([]byte(fmt.Sprintf("Task with id %d was successfully deleted.", numId)))
}
