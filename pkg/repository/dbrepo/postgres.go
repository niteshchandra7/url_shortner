package dbrepo

import (
	"context"
	"time"
)

func (m *PostgresDBRepo) Exists(link string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int
	query := `select id from links where link=$1`
	row := m.DB.QueryRowContext(ctx, query, link)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	return id > 0
}
