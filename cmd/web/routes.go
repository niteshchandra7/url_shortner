package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/niteshchandra7/url_shortner/pkg/handlers"
)

func GetRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", handlers.Home)
	return mux
}
