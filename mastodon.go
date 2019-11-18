package main

import (
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
}

func tootAll(posts []Post) {
	for _, post := range posts {

		var images = make([]io.Reader, 0, 4)

		for _, img := range post.Images {
			images = append(images, downloadMedia(img))
		}

		post.Text = html.UnescapeString(post.Text)
		if config.Mastodon.InsertEmoji {
			post.Text = insertAll(post.Text)
		}

		_, err := toot(post.Text, "", images)
		if err != nil {
			alertToOwner("tootAll(): " + err.Error() + "\n" + post.Text)
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

func toot(content, visibility string, readers []io.Reader) (st *madon.Status, err error) {

	if visibility == "" {
		visibility = "public"
	}

	var mediaIDs = make([]int64, 0, 4)

	// upload medias
	for _, reader := range readers {
		attach, err := mc.UploadMediaReader(reader, "mltd_img", "", "")
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
	}

	st, err = mc.PostStatus(status)
	return
}

func alertToOwner(content string) {
	content = config.Mastodon.Owner + " " + content
	_, err := toot(content, "direct", nil)
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
