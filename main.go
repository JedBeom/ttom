package main

import "time"

func main() {
	loadConfig()
	TwitterInit()
	MastodonInit()
	for {
		checkNew()
		time.Sleep(time.Second * 3)
	}
}
