package main

import (
	"context"
	"io"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	client     *twitter.Client
	latestPost Post
)

func TwitterInit() {
	var twitterConfig = &clientcredentials.Config{
		ClientID:     config.Twitter.ClientID,
		ClientSecret: config.Twitter.ClientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := twitterConfig.Client(context.TODO())
	client = twitter.NewClient(httpClient)
}

func getTwitterUser(acc string) (user *twitter.User, err error) {
	user, _, err = client.Users.Show(&twitter.UserShowParams{
		ScreenName: acc,
	})
	return
}

func getTweets(targetID int64, limit int) ([]twitter.Tweet, error) {
	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		UserID:    targetID,
		Count:     limit,
		TweetMode: "extended",
	})

	return tweets, err
}

func checkNew(targetID int64) {
	tweets, err := getTweets(targetID, 5)
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
		createdAt, err := tw.CreatedAtTime()
		if err != nil {
			alertToOwner("checkNew:tw.CreatedAtTime(): " + err.Error())
			continue
		}

		if createdAt.Sub(latestPost.CreatedAt).Seconds() <= 0 {
			break
		}

		newPosts = append(newPosts, tweetFilter(tw))
	}

	if len(newPosts) != 0 {
		latestPost = newPosts[0]
		go tootAll(newPosts)
	}

}

func tweetFilter(tw twitter.Tweet) (post Post) {
	post.ID = tw.ID
	post.Text = tw.FullText

	var err error
	post.CreatedAt, err = tw.CreatedAtTime()
	if err != nil {
		alertToOwner("tweetFilter(): " + err.Error())
	}

	if tw.ExtendedEntities != nil {
		for _, img := range tw.ExtendedEntities.Media {
			if img.Type == "photo" {
				post.Images = append(post.Images, img.MediaURLHttps)
			}
		}
	}

	return
}

func detectNewAvatarOrHeader(old *twitter.User) (new *twitter.User) {
	var err error
	new, err = getTwitterUser(config.Twitter.Account)
	if err != nil {
		alertToOwner("detectNewAvatarOrHeader:getTwitterUser(): " + err.Error())
		return
	}

	var avatar, header io.Reader
	if new.ProfileImageURLHttps != old.ProfileImageURLHttps {
		avatar = downloadMedia(strings.Replace(new.ProfileImageURLHttps, "_normal", "", 1))
	}

	if new.ProfileBannerURL != old.ProfileBannerURL {
		header = downloadMedia(new.ProfileBannerURL)
	}

	if avatar != nil || header != nil {
		changeProfilePhoto(avatar, header)
	}
	return
}
