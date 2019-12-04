package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ----------------------------------------------------------------------------
// *****************************          *************************************
// ***************************** Handlers *************************************
// *****************************          *************************************
// ----------------------------------------------------------------------------

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	method, url := r.FormValue("method"), r.FormValue("url")
	fmt.Printf("method: %s, url: %s\n", method, url)
	timeout := 5 * time.Second
	resp, err := SendRequest(&http.Client{
		Timeout: timeout,
	}, method, url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		log.Println(err)
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		log.Println(err)
		return
	}
	w.Write(data)
}

func AllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := GetListController().GetAll()
	data, _ := json.Marshal(tasks)
	w.Write(data)
}

func TaskIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	task, err := GetListController().GetById(numId)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(task)
	w.Write(data)
}

func RemoveTaskIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = GetListController().RemoveById(numId)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Write([]byte(fmt.Sprintf("Task with id %d was successfully deleted.", numId)))
}

func HeadersToString(reqHeader *http.Header) string {
	var result []string
	for name, headers := range *reqHeader {
		for _, h := range headers {
			result = append(result, fmt.Sprintf("%v: %v\n", name, h))
		}
	}
	return strings.Join(result, "\n")
}
