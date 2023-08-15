package handlers

import (
	"log"
	"net/http"

	"github.com/niteshchandra7/url_shortner/pkg/config"
	"github.com/niteshchandra7/url_shortner/pkg/models"
	"github.com/niteshchandra7/url_shortner/pkg/renders"
	"github.com/niteshchandra7/url_shortner/pkg/repository"
	"github.com/niteshchandra7/url_shortner/pkg/repository/dbrepo"
)

type Repository struct {
	AppConfig *config.AppConfig
	DBRepo    repository.DatabaseRepo
}

func GetNewRepo(app *config.AppConfig, repo repository.DatabaseRepo) *Repository {
	return &Repository{
		AppConfig: app,
		DBRepo:    repo,
	}
}

var Repo *Repository

func SetNewRepo(r *Repository) {
	Repo = r
}

// Home renders home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renders.New(w, r, "home.page.go.tmpl", &models.TemplateData{})
}

// Shorten shortens the posted URL
func (m *Repository) PostShorten(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	link := r.Form.Get("url")
	exists := dbrepo.NewPostgresRepo(m.AppConfig).Exists(link)
	if !exists {
		log.Println("link doesn't exist in db")
	} else {
		log.Println("link exists in db")
	}
	http.Redirect(w, r, "https://www.google.com", http.StatusSeeOther)
	//renders.New(w, r, "home.page.go.tmpl", &models.TemplateData{})

}
