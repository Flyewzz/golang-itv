package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// var err error
	PrepareConfig()
	r := GetRouter()
	ConfigureHandlers(r)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
