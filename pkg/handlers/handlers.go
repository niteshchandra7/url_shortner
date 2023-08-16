package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/niteshchandra7/url_shortner/pkg/config"
	"github.com/niteshchandra7/url_shortner/pkg/models"
	"github.com/niteshchandra7/url_shortner/pkg/renders"
	"github.com/niteshchandra7/url_shortner/pkg/repository"
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
	var shorten_link string
	shorten_link, exists := m.DBRepo.GetShortenLinkFromLink(link)
	if !exists {
		shorten_link = m.DBRepo.CreateAndInsertShortenLinkFromLink(link)
	}
	renders.New(w, r, "home.page.go.tmpl", &models.TemplateData{
		ShortenURL: os.Getenv("HOME_URL") + shorten_link,
	})
}

func (m *Repository) GetLink(w http.ResponseWriter, r *http.Request) {
	shorten_link := chi.URLParam(r, "shorten-url")
	link, exists := m.DBRepo.GetLinkFromShortenLink(shorten_link)
	if !exists {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, link, http.StatusSeeOther)
}
