package dbrepo

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// GetShortenLink gets shorten link from database
func (m *PostgresDBRepo) GetShortenLinkFromLink(link string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var shortenLink string
	var id int

	query := `select id, shorten_link from links where link=$1`

	row := m.DB.QueryRowContext(ctx, query, link)
	err := row.Scan(&id, &shortenLink)
	if err != nil {
		return shortenLink, false
	}
	return shortenLink, id > 0
}

func (m *PostgresDBRepo) CreateAndInsertShortenLinkFromLink(link string) string {
	var shorten_link string
	shorten_link, exists := m.GetShortenLinkFromLink(link)
	if exists {
		return shorten_link
	}
	shorten_link = GetMD5Hash(link)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into links (link, shorten_link) values ($1,$2)`

	m.DB.ExecContext(ctx, query, link, shorten_link)

	return shorten_link
}

// GetShortenLink gets shorten link from database
func (m *PostgresDBRepo) GetLinkFromShortenLink(shorten_link string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var link string
	var id int

	query := `select id, link from links where shorten_link=$1`

	row := m.DB.QueryRowContext(ctx, query, shorten_link)
	err := row.Scan(&id, &link)
	if err != nil {
		return link, false
	}
	return link, id > 0
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	len := len(hash)
	x := rand.Intn(len) - 8
	if x < 0 {
		x = 0
	}
	return hex.EncodeToString(hash[x : x+7])
}
