package main

import (
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

func main() {
	loadConfig()
	TwitterInit()
	MastodonInit()

	user, err := getTwitterUser(config.Twitter.Account)
	if err != nil {
		panic(err)
	}

	generateRegexp(&idolTable)

	go autoChange(user)

	for {
		checkNew(user.ID)
		time.Sleep(time.Second * time.Duration(config.Twitter.RefreshSecond))
	}
}

func autoChange(start *twitter.User) {
	var user = new(twitter.User)
	*user = *start

	for {
		user = detectNewAvatarOrHeader(user)
		time.Sleep(time.Hour * 24)
	}
}
