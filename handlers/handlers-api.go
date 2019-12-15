package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router, uh *UserHandler) {
	r.HandleFunc("/request", uh.RequestHandler).Methods("GET")
	r.HandleFunc("/requests", uh.AllRequestsHandler).Methods("GET")
	r.HandleFunc("/requests", uh.AllRequestsRemoveHandler).Methods("DELETE")
	r.HandleFunc("/requests/page/{number}", uh.PageHandler).Methods("GET")
	r.HandleFunc("/requests/{id}", uh.RequestIdHandler).Methods("GET")
	r.HandleFunc("/requests/{id}", uh.RemoveRequestIdHandler).Methods("DELETE")
}
