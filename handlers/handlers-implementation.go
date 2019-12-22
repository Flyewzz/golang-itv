package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "github.com/Flyewzz/golang-itv/features"
	"github.com/Flyewzz/golang-itv/models"
	"github.com/gorilla/mux"
)

func (hd *HandlerData) RequestHandler(w http.ResponseWriter, r *http.Request) {
	method, url := r.URL.Query().Get("method"), r.URL.Query().Get("url")
	if !CheckMethodValid(method) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	task := &models.Task{
		Method: method,
		Url:    url,
	}
	// timeout := 5 * time.Second
	// resp, err := hd.Executor.Execute(&http.Client{
	// 	Timeout: timeout,
	// }, task)
	resCh := hd.Dispatcher.AddNewTask(task)
	// result := models.NewRequest(task, resp)
	// hd.StoreController.Add(result)
	result := <-resCh
	if result.Error != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println(result.Error)
		return
	}
	data, err := json.Marshal(result.Response)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) AllRequestsHandler(w http.ResponseWriter, r *http.Request) {
	requests := hd.StoreController.GetAll()
	data, _ := json.Marshal(requests)
	w.Write(data)
}

func (hd *HandlerData) AllRequestsRemoveHandler(w http.ResponseWriter, r *http.Request) {
	hd.StoreController.RemoveAll()
	w.Write([]byte("All requests was successfully deleted."))
}

func (hd *HandlerData) PageHandler(w http.ResponseWriter, r *http.Request) {
	strNum := mux.Vars(r)["number"]
	num, err := strconv.Atoi(strNum)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	requests, err := hd.StoreController.GetByPage(num)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(requests)
	w.Write(data)
}

func (hd *HandlerData) RequestIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	request, err := hd.StoreController.GetById(numId)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(request)
	w.Write(data)
}

func (hd *HandlerData) RemoveRequestIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = hd.StoreController.RemoveById(numId)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Write([]byte(fmt.Sprintf("Request with id %d was successfully deleted.", numId)))
}
