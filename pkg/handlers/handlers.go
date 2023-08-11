package handlers

import (
	"net/http"

	"github.com/niteshchandra7/url_shortner/pkg/renders"
)

// type Repository struct {
// 	TemplateCache map[string]*template.Template
// }

// func GetNewRepo(app *config.AppConfig) *Repository {
// 	return &Repository{
// 		TemplateCache: app.TemplateCache,
// 	}
// }

// var repository *Repository

// func SetNewRepo(r *Repository) {
// 	repository = r
// }

// Home renders home page
func Home(w http.ResponseWriter, r *http.Request) {
	renders.New(w, r, "home.page.go.tmpl")
}
