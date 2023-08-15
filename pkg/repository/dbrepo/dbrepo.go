package dbrepo

import (
	"database/sql"

	"github.com/niteshchandra7/url_shortner/pkg/config"
	"github.com/niteshchandra7/url_shortner/pkg/repository"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

// NewPostgresRepo creates new postgres database repo
func NewPostgresRepo(appConfig *config.AppConfig) repository.DatabaseRepo {
	return &PostgresDBRepo{
		DB: appConfig.DB.SQL,
	}
}
