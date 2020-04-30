package main

import "time"

type Post struct {
	ID    int64
	Text  string
	Media []string

	CreatedAt time.Time
}
