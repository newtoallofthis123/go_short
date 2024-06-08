package main

import "time"

type Env struct {
	DatabaseUrl string
	Port        string
}
type CreateURLRequest struct {
	url  string
	hash string
}

type Url struct {
	hash      string
	url       string
	createdAt time.Time
}
