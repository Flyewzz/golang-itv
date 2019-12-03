package main

import (
	"sync"

	"github.com/gorilla/mux"
)

var R *mux.Router
var onceRouter sync.Once

func GetRouter() *mux.Router {
	onceRouter.Do(func() {
		R = mux.NewRouter()
	})
	return R
}
