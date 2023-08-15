package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/niteshchandra7/url_shortner/pkg/drivers"
)

// AppConfig stores app configuration
type AppConfig struct {
	Addr          string
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
	DB            *drivers.DB
}

// New creates a new app config
func New(addr string) *AppConfig {
	return &AppConfig{
		Addr:          ":8080",
		TemplateCache: map[string]*template.Template{},
	}
}

// LoadEnvironment loads all environment variable from .env file
func LoadEnvironment() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
