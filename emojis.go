package main

import (
	"regexp"
	"strings"
)

type Idol struct {
	LastName  string
	FirstName string
	EmojiName string

	Regex *regexp.Regexp
}

const (
	EmojiPrefix = ":mltd_"
	EmojiSuffix = ":"
)

var (
	Types = []string{
		"Fairy", ":mltd_fairy: Fairy",
		"Angel", ":mltd_angel: Angel",
		"Princess", ":mltd_princess: Princess",
		"SSR", ":mltd_gasha_rainbow: SSR",
	}
)

var idolTable = []Idol{
	//　All Stars
	{LastName: "아마미", FirstName: "하루카", EmojiName: "haruka"},
	{LastName: "키사라기", FirstName: "치하야", EmojiName: "chihaya"},
	{LastName: "호시이", FirstName: "미키", EmojiName: "miki"},
	{LastName: "하기와라", FirstName: "유키호", EmojiName: "yukiho"},
	{LastName: "타카츠키", FirstName: "야요이", EmojiName: "yayoi"},
	{LastName: "키쿠치", FirstName: "마코토", EmojiName: "makoto"},
	{LastName: "미나세", FirstName: "이오리", EmojiName: "iori"},
	{LastName: "시죠", FirstName: "타카네", EmojiName: "takane"},
	{LastName: "아키즈키", FirstName: "리츠코", EmojiName: "ritsuko"},
	{LastName: "미우라", FirstName: "아즈사", EmojiName: "azusa"},
	{LastName: "후타미", FirstName: "아미", EmojiName: "ami"},
	{LastName: "후타미", FirstName: "마미", EmojiName: "mami"},
	{LastName: "가나하", FirstName: "히비키", EmojiName: "hibiki"},

	// Princess
	{LastName: "카스가", FirstName: "미라이", EmojiName: "mirai"},
	{LastName: "타나카", FirstName: "코토하", EmojiName: "kotoha"},
	{LastName: "사타케", FirstName: "미나코", EmojiName: "minako"},
	{LastName: "토쿠가와", FirstName: "마츠리", EmojiName: "matsuri"},
	{LastName: "나나오", FirstName: "유리코", EmojiName: "yuriko"},
	{LastName: "타카야마", FirstName: "사요코", EmojiName: "sayoko"},
	{LastName: "마츠다", FirstName: "아리사", EmojiName: "arisa"},
	{LastName: "코사카", FirstName: "우미", EmojiName: "umi"},
	{LastName: "나카타니", FirstName: "이쿠", EmojiName: "iku"},
	{FirstName: "에밀리", EmojiName: "emily"},
	{LastName: "야부키", FirstName: "카나", EmojiName: "kana"},
	{LastName: "요코야마", FirstName: "나오", EmojiName: "nao"},
	{LastName: "후쿠다", FirstName: "노리코", EmojiName: "noriko"},

	// Fairy
	{LastName: "모가미", FirstName: "시즈카", EmojiName: "sizuka"},
	{LastName: "토코로", FirstName: "메구미", EmojiName: "megumi"},
	{FirstName: "로코", EmojiName: "roco"},
	{LastName: "텐쿠바시", FirstName: "토모카", EmojiName: "tomoka"},
	{LastName: "키타자와", FirstName: "시호", EmojiName: "shiho"},
	{LastName: "마이하마", FirstName: "아유무", EmojiName: "ayumu"},
	{LastName: "니카이도", FirstName: "치즈루", EmojiName: "chizuru"},
	{LastName: "마카베", FirstName: "미즈키", EmojiName: "mizuki"},
	{LastName: "모모세", FirstName: "리오", EmojiName: "rio"},
	{LastName: "나가요시", FirstName: "스바루", EmojiName: "subaru"},
	{LastName: "스오", FirstName: "모모코", EmojiName: "momoko"},
	{FirstName: "줄리아", EmojiName: "julia"},
	{LastName: "시라이시", FirstName: "츠무기", EmojiName: "tsumugi"},

	// Angel
	{LastName: "이부키", FirstName: "츠바사", EmojiName: "tsubasa"},
	{LastName: "시마바라", FirstName: "엘레나", EmojiName: "elena"},
	{LastName: "하코자키", FirstName: "세리카", EmojiName: "serika"},
	{LastName: "노노하라", FirstName: "아카네", EmojiName: "akane"},
	{LastName: "모치즈키", FirstName: "안나", EmojiName: "anna"},
	{LastName: "키노시타", FirstName: "히나타", EmojiName: "hinata"},
	{LastName: "바바", FirstName: "코노미", EmojiName: "konomi"},
	{LastName: "오오가미", FirstName: "타마키", EmojiName: "tamaki"},
	{LastName: "토요카와", FirstName: "후우카", EmojiName: "fuka"},
	{LastName: "미야오", FirstName: "미야", EmojiName: "miya"},
	{LastName: "시노미야", FirstName: "카렌", EmojiName: "karen"},
	{LastName: "키타카미", FirstName: "레이카", EmojiName: "reika"},
	{LastName: "사쿠라모리", FirstName: "카오리", EmojiName: "kaori"},

	// Workers
	{LastName: "오토나시", FirstName: "코토리", EmojiName: "kotori"},
	{LastName: "아오바", FirstName: "미사키", EmojiName: "misaki"},

	// ETC
	{FirstName: "SR", EmojiName: "gasha_yellow"},
}

func generateRegexp(idols *[]Idol) {
	for i := range *idols {
		var err error
		if (*idols)[i].LastName == "" {
			(*idols)[i].Regex, err = regexp.Compile(`(\A|\s|'|")(` + (*idols)[i].FirstName + `)`)
		} else {
			(*idols)[i].Regex, err = regexp.Compile(`(\A|\s|\s'|\s")(` + (*idols)[i].LastName + `\s*)?(` + (*idols)[i].FirstName + `)`)
		}

		if err != nil {
			panic(err)
		}
	}
}

func insertEmojis(text string) (output string) {
	for i := range idolTable {
		emoji := EmojiPrefix + idolTable[i].EmojiName + EmojiSuffix // :mltd_name:
		index := idolTable[i].Regex.FindStringIndex(text)           // find
		if len(index) != 0 {                                        // if exist
			if string(text[index[0]]) != " " { // if not space

				if text[index[0]] == 10 { // if newline
					emoji = "\n" + emoji + " "                           // add newline
					output = text[:index[0]] + emoji + text[index[0]+1:] // ignores old newline
				} else {
					emoji += " "
					output = text[:index[0]] + emoji + text[index[0]:] // insert
				}

			} else { // if space
				emoji = " " + emoji
				output = text[:index[0]] + emoji + text[index[0]:] // insert
			}
		}
	}

	r := strings.NewReplacer(Types...) // types replacer
	output = r.Replace(output)         // insert types emojis

	return
}
