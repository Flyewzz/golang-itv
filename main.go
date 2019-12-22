package main

import (
	"fmt"
	"net/http"

	"github.com/Flyewzz/golang-itv/handlers"
	"github.com/spf13/viper"
)

func main() {
	PrepareConfig()
	r := NewRouter()
	HandlerData := PrepareHandlerData()
	handlers.ConfigureHandlers(r, HandlerData)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
