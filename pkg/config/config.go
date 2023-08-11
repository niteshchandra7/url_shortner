package config

import "html/template"

// AppConfig stores app configuration
type AppConfig struct {
	Addr          string
	TemplateCache map[string]*template.Template
	InProduction  bool
}

// New creates a new app config
func New(addr string) *AppConfig {
	return &AppConfig{
		Addr:          ":8080",
		TemplateCache: map[string]*template.Template{},
	}
}
