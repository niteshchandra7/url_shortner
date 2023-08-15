package drivers

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

var dbConn = &DB{}

// ConnectSQL creates database pool for postgres
func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)
	dbConn.SQL = db
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return dbConn, nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DSN"))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
