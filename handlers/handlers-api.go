package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router, uh *UserHandler) {
	r.HandleFunc("/request", uh.RequestHandler).Methods("GET")
	r.HandleFunc("/tasks", uh.AllTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", uh.AllTasksRemoveHandler).Methods("DELETE")
	r.HandleFunc("/tasks/page/{number}", uh.PageTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", uh.TaskIdHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", uh.RemoveTaskIdHandler).Methods("DELETE")
}
