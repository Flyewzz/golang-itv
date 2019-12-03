package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"

	"github.com/spf13/viper"
)

var R *mux.Router
var onceRouter sync.Once

func GetRouter() *mux.Router {
	onceRouter.Do(func() {
		R = mux.NewRouter()
	})
	return R
}

func PrepareConfig() {
	viper.SetConfigFile(os.Args[1])
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read a config file: %v\n", err)
	}
}

func main() {
	// var err error
	PrepareConfig()
	// Check for connection
	r := GetRouter()
	ConfigureHandlers(r)
	// r.Use(RequireAuthentication)
	fmt.Println("Server is starting...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
