package main

import "strings"

const (
	InsertEmojis = true
)

var (
	Types = []string{
		"Fairy", ":mltd_fairy: Fairy",
		"Angel", ":mltd_angel: Angel",
		"Princess", ":mltd_princess: Princess",
	}
)

func Replace(text string) (output string) {
	replacer := strings.NewReplacer(Types...)
	output = replacer.Replace(text)
	return
}
