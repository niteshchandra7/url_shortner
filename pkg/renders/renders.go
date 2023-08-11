package renders

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/niteshchandra7/url_shortner/pkg/config"
	"github.com/niteshchandra7/url_shortner/pkg/models"
)

// Repository creates repository for render package
type Repository struct {
	TemplateCache map[string]*template.Template
	InProduction  bool
}

// GetNewRepo creates and return a new render repository
func GetNewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		TemplateCache: app.TemplateCache,
		InProduction:  app.InProduction,
	}
}

var repository *Repository

// SetNewRepo sets render repository
func SetNewRepo(r *Repository) {
	repository = r
}

// CreateTemplateCache creates a TemplateCache
func CreateTemplateCache() error {
	pages, err := filepath.Glob("./templates/*page.go.tmpl")
	if err != nil {
		log.Fatal("templates dir has no page.go.tmpl file", err)
	}
	layouts, err := filepath.Glob("./templates/*layout.go.tmpl")
	if err != nil || len(layouts) == 0 {
		log.Fatal("templates dir has no layout.go.tmpl file", err)
	}
	err = CreateTemplates(pages, layouts)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// CreateTemplates store template and pageName into TemplateCache
func CreateTemplates(pages []string, layouts []string) error {
	for _, page := range pages {
		tmpl, err := template.ParseFiles(page)
		if err != nil {
			log.Fatal(err)
		}
		pageName := strings.TrimPrefix(page, "templates/")
		tmpl = tmpl.New(pageName)
		tmpl.ParseFiles(layouts...)
		repository.TemplateCache[pageName] = tmpl
	}
	return nil
}

// New renders the page from TemplateCache
func New(w http.ResponseWriter, r *http.Request, pageName string) {
	if !repository.InProduction {
		layouts, err := filepath.Glob("./templates/*layout.go.tmpl")
		if err != nil || len(layouts) == 0 {
			log.Fatal("templates dir has no layout.go.tmpl file", err)
		}
		CreateTemplates([]string{"templates/" + pageName}, layouts)
	}
	tmpl, ok := repository.TemplateCache[pageName]
	if !ok {
		log.Fatal("page not found in template cache")

	}
	tmpl.Execute(w, models.TemplateData{})
}
