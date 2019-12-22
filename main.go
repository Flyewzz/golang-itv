package main

import (
	"fmt"
	"net/http"

	"github.com/Flyewzz/golang-itv/executor"
	"github.com/Flyewzz/golang-itv/handlers"
	"github.com/Flyewzz/golang-itv/store"
	"github.com/Flyewzz/golang-itv/workers/dispatcher"
	"github.com/spf13/viper"
)

func main() {
	// var err error
	PrepareConfig()
	r := NewRouter()
	storeController := store.NewStoreController(viper.GetInt("itemsPerPage"), 0)
	executor := executor.NewHttpExecutor()
	countWorkers := viper.GetInt("workers.count")
	maxTasks := viper.GetInt("tasks.max")
	dispatcher := dispatcher.NewDispatcher(countWorkers, maxTasks, executor, storeController)
	dispatcher.Dispatch()
	HandlerData := handlers.NewHandlerData(executor, storeController, dispatcher)
	handlers.ConfigureHandlers(r, HandlerData)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
