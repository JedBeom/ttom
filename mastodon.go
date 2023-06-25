package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/JedBeom/madon/v3"
)

var (
	mc *madon.Client
)

func MastodonInit() {
	token := madon.UserToken{
		AccessToken: config.Mastodon.AccessToken,
		CreatedAt:   time.Now().UnixNano(),
		Scope:       "read write",
		TokenType:   "urn:ietf:wg:oauth:2.0:oob",
	}

	var err error
	mc, err = madon.RestoreApp("Twitter", config.Mastodon.Instance,
		config.Mastodon.ClientKey,
		config.Mastodon.ClientSecret,
		&token)

	if err != nil {
		panic(err)
	}

	alertToOwner("가동을 시작합니다")
}

func tootPosts(posts []Post) {
	for i := range posts {
		post, err := tootPost(posts[i], "")
		posts[i] = post
		if err != nil {
			alertToOwner("tootPosts(): " + err.Error() + "\n" + posts[i].Content)
		}
	}

	for _, post := range posts {
		if post.TootID == 0 {
			continue
		}

		err := insertPost(post)
		if err != nil {
			alertToOwner("insertPost(): " + err.Error() + "\n" + fmt.Sprintf("%+v", post))
		}
	}
}

func downloadMedia(link string) io.Reader {
	resp, err := http.Get(link)
	if err != nil {
		alertToOwner("downloadMedia(): " + err.Error())
		return nil
	}

	return resp.Body
}

func tootPost(post Post, visibility string) (Post, error) {
	var inReplyTo int64

	if post.IsBoosted || post.IsQuoted || post.IsReplied {
		var err error
		post.SubjectTootID, err = getTootIDByTweetID(post.SubjectTweetID)
		if err != nil {
			log.Println("getTootIDByTweetID():", err)
		}
	}

	if post.IsBoosted {
		// If the tweet was for retweet, follow below cases.
		// 1) The original tweet is in the table:
		// 		Boost the toot that matches the original tweet.
		// 2) THe original tweet is NOT in the table:
		// 		Do not take any action. End this function.
		if post.SubjectTootID == 0 {
			return post, nil
		}

		err := mc.ReblogStatus(post.SubjectTootID)
		return post, err

	} else if post.IsQuoted {
		// If the tweet was quoted tweet, follow below cases.
		// 1) The original tweet is in the table:
		// 		Get the url of the original toot matches the original tweet.
		// 		Concat next to the content(e.g. content\n\nBT: planet.moe/...)
		// 2) NOT in the table:
		// 		Make the url of the original tweet.
		// 		Concat next to the content(e.g. cnotent\n\nBT: twitter.com/...)
		url := fmt.Sprintf("https://twitter.com/%s/status/%s", config.Twitter.Account, post.SubjectTweetID)
		if post.SubjectTootID != 0 {
			st, err := mc.GetStatus(post.SubjectTootID)
			if err == nil {
				url = st.URI
			}
		}

		post.Content = "Quote: " + url + "\n\n" + post.Content

	} else if post.IsReplied {
		// If the tweet was a reply, follow below cases.
		// 1) The original tweet is in the table:
		// 		Reply to the original toot.
		// 2) NOT in the table:
		// 		Concat before the content(e.g. Replying to URL\n\ncontent)
		inReplyTo = post.SubjectTootID

		if post.SubjectTootID == 0 {
			replyURL := fmt.Sprintf("https://twitter.com/%s/status/%s", config.Twitter.Account, post.SubjectTweetID)
			post.Content = "Replying to " + replyURL + "\n\n" + post.Content
		}
	}

	var readers = make([]io.Reader, 0, 4)

	for _, img := range post.Media {
		readers = append(readers, downloadMedia(img))
	}

	post.Content = html.UnescapeString(post.Content)
	if config.Mastodon.InsertEmoji {
		post.Content = insertEmoji(post.Content)
	}

	st, err := toot(post.Content, "", inReplyTo, readers)
	if st != nil {
		post.TootID = st.ID
	}
	return post, err
}

func toot(content, visibility string, inReplyTo int64, readers []io.Reader) (st *madon.Status, err error) {
	if visibility == "" {
		visibility = config.Mastodon.Visibility
	}

	var mediaIDs = make([]int64, 0, 4)

	// upload medias
	for _, reader := range readers {
		attach, err := mc.UploadMediaReader(reader, "mltd_media", "© BANDAI NAMCO", "")
		if err != nil {
			alertToOwner("toot(): " + err.Error())
			continue
		}

		mediaIDs = append(mediaIDs, attach.ID)
	}

	status := madon.PostStatusParams{
		Text:       content,
		Visibility: visibility,
		MediaIDs:   mediaIDs,
		InReplyTo:  inReplyTo,
	}

	st, err = mc.PostStatus(status)
	return
}

func alertToOwner(content string) {
	content = config.Mastodon.Owner + " " + content
	_, err := toot(content, "direct", 0, nil)
	if err != nil {
		log.Println(err)
	}
}

func changeProfilePhoto(avatar, header io.Reader) {
	_, err := mc.UpdateAccountReader(madon.UpdateAccountParams{
		AvatarImageReader: avatar,
		HeaderImageReader: header,
	})

	if err != nil {
		alertToOwner("changeProfilePhoto:UpdateAccount(): " + err.Error())
	}
}
