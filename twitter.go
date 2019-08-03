package main

import (
	"context"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	client           *twitter.Client
	latestPost       Post
	targetTwitterAcc *twitter.User
)

func TwitterInit() {
	var twitterConfig = &clientcredentials.Config{
		ClientID:     config.Twitter.ClientID,
		ClientSecret: config.Twitter.ClientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := twitterConfig.Client(context.TODO())
	client = twitter.NewClient(httpClient)

	var err error
	targetTwitterAcc, _, err = client.Users.Show(&twitter.UserShowParams{
		ScreenName: config.Twitter.Account,
	})
	if err != nil {
		panic(err)
	}
}

func getTweets(limit int) ([]twitter.Tweet, error) {
	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		UserID:    targetTwitterAcc.ID,
		Count:     limit,
		TweetMode: "extended",
	})

	return tweets, err
}

func checkNew() {
	tweets, err := getTweets(5)
	if err != nil {
		alertToOwner("checkNew(): " + err.Error())
		return
	}

	if len(tweets) == 0 {
		return
	}

	if latestPost.ID == 0 {
		latestPost = tweetFilter(tweets[0])
		return
	}

	var newPosts []Post
	for _, tw := range tweets {
		if latestPost.ID == tw.ID {
			break
		}

		newPosts = append(newPosts, tweetFilter(tw))
	}

	if len(newPosts) >= 5 {
		alertToOwner("It looks some tweets were deleted.")
		latestPost = newPosts[0]
		return
	}

	if len(newPosts) != 0 {
		latestPost = newPosts[0]
		go tootAll(newPosts)
	}

}

func tweetFilter(tw twitter.Tweet) (post Post) {
	post.ID = tw.ID
	post.Text = tw.FullText

	if tw.ExtendedEntities != nil {
		for _, img := range tw.ExtendedEntities.Media {
			if img.Type == "photo" {
				post.Images = append(post.Images, img.MediaURLHttps)
			}
		}
	}

	return
}
