package main

import (
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router) {

	r.HandleFunc("/", MainFunc).Methods("GET")

	// Configure the static directory
	// r.PathPrefix("/").Handler(http.FileServer(http.Dir(viper.GetString("static_path"))))
}
