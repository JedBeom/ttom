package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func initSQL() {
	var err error

	db, err = sql.Open("sqlite3", config.SQL.Filename)
	if err != nil {
		panic(err)
	}
	createTable()

	latestTweetDate, err = getLastPostDate()
	if err != nil {
		log.Println("Error on getLastPostDate():", err, "\nparsed:", latestTweetDate)
	}
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS posts (
		toot_id INTEGER NOT NULL PRIMARY KEY,
		tweet_id TEXT,
		is_boosted INTEGER,
		is_quoted INTEGER,
		is_replied INTEGER,
		subject_toot_id INTEGER,
		subject_tweet_id TEXT,
		content TEXT,
		created_at INTEGER
	)`

	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}

func insertPost(post Post) error {
	query, err := db.Prepare(`INSERT INTO posts VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = query.Exec(post.TootID, post.TweetID, post.IsBoosted, post.IsQuoted, post.IsReplied, post.SubjectTootID, post.SubjectTweetID, post.Content, post.CreatedAt)
	return err
}

func getTootIDByTweetID(tweetID string) (int64, error) {
	query, err := db.Prepare(`SELECT toot_id FROM posts WHERE tweet_id = ? LIMIT 1`)
	if err != nil {
		return 0, err
	}

	row := query.QueryRow(tweetID)
	if row.Err() != nil {
		return 0, row.Err()
	}

	var tootID int64
	err = row.Scan(&tootID)
	return tootID, err
}

func getLastPostDate() (time.Time, error) {
	date := time.Now()
	row := db.QueryRow(`SELECT created_at FROM posts ORDER BY created_at DESC`)
	if row.Err() != nil {
		return date, row.Err()
	}

	var timestamp string
	err := row.Scan(&timestamp)
	if err != nil {
		return date, err
	}

	return time.Parse("2006-01-02 15:04:05", timestamp[:len(timestamp)-6])
}
