package main

import (
	"testing"
)

func Test_insertEmoji(t *testing.T) {
	generateRegexp(&idolTable)
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"event",
			args{`イベント『プラチナスターシアター』開演です！
イベント楽曲『Beat the World!!!』や、イベント限定コミュ(全7話！)も楽しんでくださいね♪

【イベント限定カード】
衣装付きSR 菊地真
SR 舞浜歩
#ミリシタ`},
			`イベント『​:mltd_pst:​プラチナスターシアター』開演です！
イベント楽曲『Beat the World!!!』や、イベント限定コミュ(全7話！)も楽しんでくださいね♪

【イベント限定カード】
衣装付き​:mltd_gasha_yellow:​SR ​:mltd_makoto:​菊地真
​:mltd_gasha_yellow:​SR ​:mltd_ayumu:​舞浜歩
#ミリシタ`,
		},
		{
			"SSR CARD",
			args{`今回追加されるのり子ちゃんのカードには
新衣装『ミルクティーン・ダイヤ』がついてきます！
この機会に是非ゲットしてください♪

のり子「この服……やっぱり、変、だよね？
店員さんはほめてくれたけど、なんか
落ち着かなくって。大丈夫……かな……？」
#ミリシタ`},
			`今回追加される​:mltd_noriko:​のり子ちゃんのカードには
新衣装『ミルクティーン・ダイヤ』がついてきます！
この機会に是非ゲットしてください♪

​:mltd_noriko:​のり子「この服……やっぱり、変、だよね？
店員さんはほめてくれたけど、なんか
落ち着かなくって。大丈夫……かな……？」
#ミリシタ`,
		},
		{
			"BIRTHDAY",
			args{`本日はロコちゃんの誕生日です！
劇場のホワイトボードに、風花さんと千鶴さんと茜ちゃんと瑞希ちゃんと雪歩ちゃんと歩ちゃんがメッセージを書いてくれたみたいですよ♪
プロデューサーさんもお祝いパーティに来てくださいね！
#ミリシタ`},
			`本日は​:mltd_roco:​ロコちゃんの誕生日です！
劇場のホワイトボードに、​:mltd_fuka:​風花さんと千鶴さんと​:mltd_akane:​茜ちゃんと​:mltd_mizuki:​瑞希ちゃんと​:mltd_yukiho:​雪歩ちゃんと​:mltd_ayumu:​歩ちゃんがメッセージを書いてくれたみたいですよ♪
プロデューサーさんもお祝いパーティに来てくださいね！
#ミリシタ`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertEmoji(tt.args.text); got != tt.want {
				t.Errorf("insertEmoji() = %v\n\n===\n\n want %v", got, tt.want)
			}
		})
	}
}
