package main

import "time"

type Post struct {
	TootID  int64
	TweetID string

	IsBoosted bool
	IsQuoted  bool
	IsReplied bool

	SubjectTootID  int64
	SubjectTweetID string

	Content string
	Media   []string

	CreatedAt time.Time
}
