package main

import (
	"log"
	"os"

	"github.com/Flyewzz/golang-itv/executor"
	"github.com/Flyewzz/golang-itv/handlers"
	"github.com/Flyewzz/golang-itv/store"
	"github.com/Flyewzz/golang-itv/workers/dispatcher"
	"github.com/spf13/viper"
)

func PrepareConfig() {
	viper.SetConfigFile(os.Args[1])
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read a config file: %v\n", err)
	}
}

func PrepareHandlerData() *handlers.HandlerData {
	storeController := store.NewStoreController(viper.GetInt("itemsPerPage"), 0)
	executor := executor.NewHttpExecutor()
	countWorkers := viper.GetInt("workers.count")
	maxTasks := viper.GetInt("tasks.max")
	workersTimeout := viper.GetInt("workers.timeout")
	dispatcher := dispatcher.NewDispatcher(countWorkers, maxTasks, workersTimeout, executor, storeController)
	dispatcher.Dispatch()
	return handlers.NewHandlerData(executor, storeController, dispatcher)
}
