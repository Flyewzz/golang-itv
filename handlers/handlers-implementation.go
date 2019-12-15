package handlers

import (
	"encoding/json"
	"fmt"
	. "github.com/Flyewzz/golang-itv/features"
	"github.com/Flyewzz/golang-itv/models"
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
	task := &models.Task{
		Method: method,
		Url:    url,
	}
	timeout := 5 * time.Second
	resp, err := uh.Executor.Execute(&http.Client{
		Timeout: timeout,
	}, task)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	result := models.NewRequest(task, resp)
	uh.StoreController.Add(result)
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(data)
}

func (uh *UserHandler) AllRequestsHandler(w http.ResponseWriter, r *http.Request) {
	requests := uh.StoreController.GetAll()
	data, _ := json.Marshal(requests)
	w.Write(data)
}

func (uh *UserHandler) AllRequestsRemoveHandler(w http.ResponseWriter, r *http.Request) {
	uh.StoreController.RemoveAll()
	w.Write([]byte("All requests was successfully deleted."))
}

func (uh *UserHandler) PageHandler(w http.ResponseWriter, r *http.Request) {
	strNum := mux.Vars(r)["number"]
	num, err := strconv.Atoi(strNum)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	requests, err := uh.StoreController.GetByPage(num)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(requests)
	w.Write(data)
}

func (uh *UserHandler) RequestIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	request, err := uh.StoreController.GetById(numId)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(request)
	w.Write(data)
}

func (uh *UserHandler) RemoveRequestIdHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte(fmt.Sprintf("Request with id %d was successfully deleted.", numId)))
}
