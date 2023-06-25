package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Mastodon struct {
		Instance     string `json:"instance"`
		ClientKey    string `json:"client_key"`
		ClientSecret string `json:"client_secret"`
		AccessToken  string `json:"access_token"`

		InsertEmoji bool   `json:"insert_emoji"`
		Owner       string `json:"owner"`
		Visibility  string `json:"visibility"`
	} `json:"mastodon"`

	Twitter struct {
		Account            string `json:"account"`
		TweetFetchLimit    int    `json:"tweet_fetch_limit"`
		TweetRefreshSecond int    `json:"tweet_refresh_sec"`
		ProfileRefreshHour int    `json:"profile_refresh_hour"`
	} `json:"twitter"`

	SQL struct {
		Filename string `json:"filename"`
	} `json:"sql"`
}

var config Config

func loadConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
}
