package repository

import (
	"URLShortener/internal/model"
	"database/sql"

	_ "github.com/lib/pq"
)

type ShortenerRepository struct {
	db *sql.DB
}

func NewShortenerRepository(conn *sql.DB) *ShortenerRepository {
	return &ShortenerRepository{db: conn}
}

func (r ShortenerRepository) CreateShortLink(link string, shortLink string) (*model.Link, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO short_url(long_value, short_value) VALUES ($1, $2) RETURNING id, long_value, short_value", link, shortLink,
	).Scan(
		&id,
		&link,
		&shortLink,
	)

	if err != nil {
		return &model.Link{}, err
	}
	return model.NewLink(int64(id), link, shortLink), nil
}

func (r ShortenerRepository) GetLinkFromShort(shortLink string) (*model.Link, error) {
	rows := r.db.QueryRow(
		"SELECT id, long_value, short_value FROM short_url WHERE short_value = $1", shortLink,
	)

	var link string
	var shortValue string
	var id int64

	err := rows.Scan(&id, &link, &shortValue)
	if err != nil {
		return &model.Link{}, err
	}

	return model.NewLink(id, link, shortValue), nil
}
