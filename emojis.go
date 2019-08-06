package main

import (
	"strings"
)

type Idol struct {
	LastName  string
	FirstName string
	EmojiName string
}

const (
	InsertEmojis = true
	EmojiPrefix  = ":mltd_"
	EmojiSuffix  = ":"
)

var (
	Types = []string{
		"Fairy", ":mltd_fairy: Fairy",
		"Angel", ":mltd_angel: Angel",
		"Princess", ":mltd_princess: Princess",
	}
)

var IdolTable = []Idol{
	//　All Stars
	{"아마미", "하루카", "haruka"},
	{"키사라기", "치하야", "chihaya"},
	{"호시이", "미키", "miki"},
	{"하기와라", "유키호", "yukiho"},
	{"타카츠키", "야요이", "yayoi"},
	{"키쿠치", "마코토", "makoto"},
	{"미나세", "이오리", "iori"},
	{"시죠", "타카네", "takane"},
	{"아키즈키", "리츠코", "ritsuko"},
	{"미우라", "아즈사", "azusa"},
	{"후타미", "아미", "ami"},
	{"후타미", "마미", "mami"},
	{"가나하", "히비키", "hibiki"},

	// Princess
	{"카스가", "미라이", "mirai"},
	{"타나카", "코토하", "kotoha"},
	{"사타케", "미나코", "minako"},
	{"토쿠가와", "마츠리", "matsuri"},
	{"나나오", "유리코", "yuriko"},
	{"타카야마", "사요코", "sayoko"},
	{"마츠다", "아리사", "arisa"},
	{"코사카", "우미", "umi"},
	{"나카타니", "이쿠", "iku"},
	{"에밀리", "스튜어트", "emily"},
	{"야부키", "카나", "kana"},
	{"요코야마", "나오", "nao"},
	{"후쿠다", "노리코", "noriko"},

	// Fairy
	{"모가미", "시즈카", "sizuka"},
	{"토코로", "메구미", "megumi"},
	{"", "로코", "roco"},
	{"텐쿠바시", "토모카", "tomoka"},
	{"키타자와", "시호", "shiho"},
	{"마이하마", "아유무", "ayumu"},
	{"니카이도", "치즈루", "chizuru"},
	{"마카베", "미즈키", "mizuki"},
	{"모모세", "리오", "rio"},
	{"나가요시", "스바루", "subaru"},
	{"스오", "모모코", "momoko"},
	{"", "줄리아", "julia"},
	{"시라이시", "츠무기", "tsumugi"},

	// Angel
	{"이부키", "츠바사", "tsubasa"},
	{"시마바라", "엘레나", "elena"},
	{"하코자키", "세리카", "serika"},
	{"노노하라", "아카네", "akane"},
	{"모치즈키", "안나", "anna"},
	{"키노시타", "히나타", "hinata"},
	{"바바", "코노미", "konomi"},
	{"오오가미", "타마키", "tamaki"},
	{"토요카와", "후우카", "fuka"},
	{"미야오", "미야", "miya"},
	{"시노미야", "카렌", "karen"},
	{"키타카미", "레이카", "reika"},
	{"사쿠라모리", "카오리", "kaori"},

	// Workers
	{"오토나시", "코토리", "kotori"},
	{"아오바", "미사키", "misaki"},
}

var gashaTable = []string{
	"SSR", EmojiPrefix + "gasha_rainbow" + EmojiSuffix + " SSR",
	// " SR", " " + EmojiPrefix + "gasha_yellow" + EmojiSuffix + " SR",
}

func Replace(text string) (output string) {
	tables := append(Types, gashaTable...)
	replacer := strings.NewReplacer(tables...)
	output = replacer.Replace(text)

	for _, idol := range IdolTable {
		emoji := EmojiPrefix + idol.EmojiName + EmojiSuffix
		if idol.LastName != "" && strings.Contains(output, idol.LastName) {
			output = strings.Replace(output, idol.LastName, emoji+" "+idol.LastName, 1)
		} else if strings.Contains(output, idol.FirstName) {
			output = strings.Replace(output, idol.FirstName, emoji+" "+idol.FirstName, 1)
		}
	}
	return
}
