package main

import (
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	loadConfig()
	TwitterInit()
	MastodonInit()
	initSQL()

	user, err := getTwitterUser(config.Twitter.Account)
	if err != nil {
		panic(err)
	}

	generateRegexp(&idolTable)

	go autoChange(user)

	for {
		checkNew(user.UserID)
		time.Sleep(time.Second * time.Duration(config.Twitter.TweetRefreshSecond))
	}
}

func autoChange(start twitterscraper.Profile) {
	// user := start
	var user twitterscraper.Profile

	for {
		user = detectNewAvatarOrHeader(user)
		time.Sleep(time.Hour * time.Duration(config.Twitter.ProfileRefreshHour))
	}
}
