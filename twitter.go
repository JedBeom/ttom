package main

import (
	"context"
	"io"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	tc         *twitter.Client
	latestPost Post
)

func TwitterInit() {
	var twitterConfig = &clientcredentials.Config{
		ClientID:     config.Twitter.ClientID,
		ClientSecret: config.Twitter.ClientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := twitterConfig.Client(context.TODO())
	tc = twitter.NewClient(httpClient)
}

// User를 불러옵니다
func getTwitterUser(acc string) (user *twitter.User, err error) {
	user, _, err = tc.Users.Show(&twitter.UserShowParams{
		ScreenName: acc,
	})
	return
}

// 트윗을 불러옵니다
func getTweets(targetID int64, limit int) ([]twitter.Tweet, error) {
	tweets, _, err := tc.Timelines.UserTimeline(&twitter.UserTimelineParams{
		UserID:    targetID,
		Count:     limit,
		TweetMode: "extended",
	})

	return tweets, err
}

// 새 트윗이 있는지 체크
func checkNew(targetID int64) {
	// 5개 트윗 불러오기
	tweets, err := getTweets(targetID, 5)
	if err != nil {
		return
	}

	// 아무 트윗도 안불러왔으면 끝
	if len(tweets) == 0 {
		return
	}

	// cached tweet이 없다면
	if latestPost.ID == 0 {
		latestPost = tweetFilter(tweets[0])
		return
	}

	var newPosts []Post
	for _, tw := range tweets {
		createdAt, err := tw.CreatedAtTime()
		if err != nil {
			continue
		}

		// 만약 캐시된 트윗보다 더 오래전의 트윗이라면 브레이크
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

// Tweet -> Post
func tweetFilter(tw twitter.Tweet) (post Post) {
	// ID, Text Get
	post.ID = tw.ID
	post.Text = tw.FullText

	post.CreatedAt, _ = tw.CreatedAtTime()

	// Image가 있다면
	if tw.ExtendedEntities != nil {
		for _, img := range tw.ExtendedEntities.Media {
			if img.Type == "photo" {
				post.Images = append(post.Images, img.MediaURLHttps)
			}
		}
	}

	return
}

// 새로운 아바타나 헤더가 있는지 확인
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

	// 만약 있다면 교체
	if avatar != nil || header != nil {
		changeProfilePhoto(avatar, header)
	}
	return
}
