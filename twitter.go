package main

import (
	"io"
	"strings"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

var (
	tc              *twitterscraper.Scraper
	latestTweetDate time.Time
)

func TwitterInit() {
	tc = twitterscraper.New()
}

// User를 불러옵니다
func getTwitterUser(acc string) (user twitterscraper.Profile, err error) {
	user, err = tc.GetProfile(acc)
	return
}

// 트윗을 불러옵니다
func getTweets(userID string, limit int) ([]*twitterscraper.Tweet, error) {
	tweets, _, err := tc.FetchTweetsByUserID(userID, limit, "")

	return tweets, err
}

// 새 트윗이 있는지 체크
func checkNew(userID string) {
	// 5개 트윗 불러오기
	tweets, err := getTweets(userID, 5)
	if err != nil {
		return
	}

	// 아무 트윗도 안불러왔으면 끝
	if len(tweets) == 0 {
		return
	}

	var newPosts []Post
	for _, tw := range tweets {
		// 만약 캐시된 트윗보다 더 오래전의 트윗이라면 브레이크
		if tw.TimeParsed.Sub(latestTweetDate).Seconds() <= 0 {
			break
		}

		newPosts = append(newPosts, tweetFilter(*tw))
	}

	if len(newPosts) != 0 {
		latestTweetDate = tweets[0].TimeParsed
		go tootPosts(newPosts)
	}

}

// Tweet -> Post
func tweetFilter(tw twitterscraper.Tweet) (post Post) {
	post.TweetID = tw.ID
	post.Content = tw.Text

	if tw.IsRetweet {
		post.IsBoosted = true
		post.SubjectTweetID = tw.RetweetedStatusID
	} else if tw.IsQuoted {
		post.IsQuoted = true
		post.SubjectTweetID = tw.QuotedStatusID
	} else if tw.IsReply {
		post.IsReplied = true
		post.SubjectTweetID = tw.InReplyToStatusID
	}

	post.Media = make([]string, 0, 4)
	for _, photo := range tw.Photos {
		post.Media = append(post.Media, photo.URL)
	}
	for _, video := range tw.Videos {
		post.Media = append(post.Media, video.URL)
	}

	post.CreatedAt = tw.TimeParsed
	return
}

// 새로운 아바타나 헤더가 있는지 확인
func detectNewAvatarOrHeader(old twitterscraper.Profile) (new twitterscraper.Profile) {
	var err error
	new, err = getTwitterUser(config.Twitter.Account)
	if err != nil {
		alertToOwner("detectNewAvatarOrHeader:getTwitterUser(): " + err.Error())
		return
	}

	var avatar, header io.Reader
	if new.Avatar != old.Avatar {
		avatar = downloadMedia(strings.Replace(new.Avatar, "_normal", "", 1))
	}

	if new.Banner != old.Banner {
		header = downloadMedia(new.Banner)
	}

	// 만약 있다면 교체
	if avatar != nil || header != nil {
		changeProfilePhoto(avatar, header)
	}
	return
}
