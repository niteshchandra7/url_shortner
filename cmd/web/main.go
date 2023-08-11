package main

import (
	"fmt"
	"net/http"

	"github.com/niteshchandra7/url_shortner/pkg/config"
	"github.com/niteshchandra7/url_shortner/pkg/renders"
)

const (
	portNumber = ":8080"
)

var appConfig *config.AppConfig

func main() {
	appConfig = config.New(portNumber)
	setupApplication()
	server := &http.Server{
		Addr:    appConfig.Addr,
		Handler: GetRoutes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func setupApplication() {
	renders.SetNewRepo(renders.GetNewRepo(appConfig))
	renders.CreateTemplateCache()
}
