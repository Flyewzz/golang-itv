package main

import (
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router) {
	r.HandleFunc("/", MainFunc).Methods("GET")
}