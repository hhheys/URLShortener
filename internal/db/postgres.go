package db

import (
	"database/sql"
	"log"
)

func CreateConnection(connURL string) *sql.DB {
	conn, err := sql.Open("postgres", connURL)
	if err != nil {
		log.Fatalf("couldnt connect to database %e", err)
	}
	return conn
}
