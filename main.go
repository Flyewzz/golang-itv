package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// var err error
	PrepareConfig()
	// Check for connection
	go FetchTaskWorker()
	r := GetRouter()
	ConfigureHandlers(r)
	fmt.Println("Server is starting...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
