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
	} `json:"mastodon"`

	Twitter struct {
		ClientID      string `json:"client_id"`
		ClientSecret  string `json:"client_secret"`
		Account       string `json:"account"`
		RefreshSecond int    `json:"refresh_sec"`
	} `json:"twitter"`
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
