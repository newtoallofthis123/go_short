package main

import "time"

type Env struct {
	DatabaseUrl string `json:"database_url,omitempty"`
	Port        string `json:"port,omitempty"`
}
type CreateURLRequest struct {
	Url  string `json:"url,omitempty"`
	Hash string `json:"hash,omitempty"`
}

type Url struct {
	Hash      string    `json:"hash,omitempty"`
	Url       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
