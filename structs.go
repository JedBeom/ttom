package main

import "time"

type Post struct {
	ID     int64
	Text   string
	Images []string

	CreatedAt time.Time
}
