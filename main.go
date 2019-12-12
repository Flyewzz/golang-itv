package main

import (
	"fmt"
	"net/http"

	"github.com/Flyewzz/golang-itv/executor"
	"github.com/Flyewzz/golang-itv/handlers"
	"github.com/Flyewzz/golang-itv/store"
	"github.com/spf13/viper"
)

func main() {
	// var err error
	PrepareConfig()
	r := NewRouter()
	storeController := store.NewStoreController(viper.GetInt("itemsPerPage"), 0)
	ex := executor.NewHttpExecutor()
	userHandlers := handlers.NewUserHandler(ex, storeController)
	handlers.ConfigureHandlers(r, userHandlers)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
