package main

import (
	"fmt"
	"net/http"

	"github.com/niteshchandra7/url_shortner/pkg/config"
)

var appConfig config.AppConfig

func main() {
	appConfig.Addr = ":8080"
	server := &http.Server{
		Addr:    appConfig.Addr,
		Handler: GetRoutes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
