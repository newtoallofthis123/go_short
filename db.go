package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Store interface {
	Init() error
	InsertURL(req CreateURLRequest) error
	GetUrlByHash(hash string) (string, error)
	GetUrlInfoByHash(hash string) (Url, error)
}

type DbInstance struct {
	db *sql.DB
}

func NewDBInstance() *DbInstance {
	env := GetEnv()
	db, err := sql.Open("postgres", env.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	return &DbInstance{
		db: db,
	}
}

func (pq *DbInstance) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		hash VARCHAR(8) PRIMARY KEY,
		url TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	)
	`

	_, err := pq.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

type HashExistsError struct {
	msg string
}

func (e *HashExistsError) Error() string {
	return e.msg
}

func (pq *DbInstance) InsertURL(req CreateURLRequest) error {
	existing, err := pq.GetUrlByHash(req.Hash)
	if err == nil && existing != "" {
		return &HashExistsError{"Hash already exists"}
	}

	query := `
	INSERT INTO urls (hash, url)
	VALUES ($1, $2)
	`

	_, err = pq.db.Exec(query, req.Hash, req.Url)

	return err
}

func (pq *DbInstance) GetUrlByHash(hash string) (string, error) {
	query := `
	SELECT * FROM urls WHERE hash = $1
	`

	var url Url

	row := pq.db.QueryRow(query, hash)
	err := row.Scan(&url.Hash, &url.Url, &url.CreatedAt)
	if err != nil {
		return "", err
	}

	return url.Url, nil
}

func (pq *DbInstance) GetUrlInfoByHash(hash string) (Url, error) {
	query := `
	SELECT * FROM urls WHERE hash = $1
	`

	var url Url

	row := pq.db.QueryRow(query, hash)
	err := row.Scan(&url.Hash, &url.Url, &url.CreatedAt)
	if err != nil {
		return url, err
	}

	return url, nil
}
