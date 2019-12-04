package main

import (
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router) {
	r.HandleFunc("/request", RequestHandler).Methods("GET")
	r.HandleFunc("/tasks", AllTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", TaskIdHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", RemoveTaskIdHandler).Methods("DELETE")
}
