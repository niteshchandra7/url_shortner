package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/niteshchandra7/url_shortner/pkg/handlers"
)

// GetRoutes return a mux to server
func GetRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Get("/", handlers.Repo.Home)
	mux.Post("/shorten", handlers.Repo.PostShorten)
	return mux
}
