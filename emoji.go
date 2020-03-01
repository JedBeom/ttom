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
	ZeroWidth   = "​"
	EmojiPrefix = ZeroWidth + ":mltd_"
	EmojiSuffix = ":" + ZeroWidth
)

var (
	EasyIcons = []string{
		"765PRO", ZeroWidth + ":765pro:" + ZeroWidth + "765PRO",
		"FAIRY", emojinize("fairy") + "FAIRY",
		"ANGEL", emojinize("angel") + "ANGEL",
		"PRINCESS", emojinize("princess") + "PRINCESS",
		"SSR", emojinize("gasha_rainbow") + "SSR",
		"プラチナスター", emojinize("pst") + "プラチナスター",
		"ミリオンフェス", emojinize("fes") + "ミリオンフェス",
		"ミリオンジュエル", emojinize("jewel") + "ミリオンジュエル",
	}
)

var idolTable = []Idol{
	//　All Stars
	{LastName: "天海", FirstName: "春香", EmojiName: "haruka"},
	{LastName: "如月", FirstName: "千早", EmojiName: "chihaya"},
	{LastName: "星井", FirstName: "美希", EmojiName: "miki"},
	{LastName: "萩原", FirstName: "雪歩", EmojiName: "yukiho"},
	{LastName: "高槻", FirstName: "やよい", EmojiName: "yayoi"},
	{LastName: "菊地", FirstName: "真", EmojiName: "makoto"},
	{LastName: "水瀬", FirstName: "伊織", EmojiName: "iori"},
	{LastName: "四条", FirstName: "貴音", EmojiName: "takane"},
	{LastName: "秋月", FirstName: "律子", EmojiName: "ritsuko"},
	{LastName: "三浦", FirstName: "あずさ", EmojiName: "azusa"},
	{LastName: "双海", FirstName: "亜美", EmojiName: "ami"},
	{LastName: "双海", FirstName: "真美", EmojiName: "mami"},
	{LastName: "我那覇", FirstName: "響", EmojiName: "hibiki"},

	// Princess
	{LastName: "春日", FirstName: "未来", EmojiName: "mirai"},
	{LastName: "田中", FirstName: "琴葉", EmojiName: "kotoha"},
	{LastName: "佐竹", FirstName: "美奈子", EmojiName: "minako"},
	{LastName: "徳川", FirstName: "まつり", EmojiName: "matsuri"},
	{LastName: "七尾", FirstName: "百合子", EmojiName: "yuriko"},
	{LastName: "高山", FirstName: "紗代子", EmojiName: "sayoko"},
	{LastName: "松田", FirstName: "亜利沙", EmojiName: "arisa"},
	{LastName: "高坂", FirstName: "海美", EmojiName: "umi"},
	{LastName: "中谷", FirstName: "育", EmojiName: "iku"},
	{FirstName: "エミリー", EmojiName: "emily"},
	{LastName: "矢吹", FirstName: "可奈", EmojiName: "kana"},
	{LastName: "横山", FirstName: "奈緒", EmojiName: "nao"},
	{LastName: "福田", FirstName: "のり子", EmojiName: "noriko"},

	// Fairy
	{LastName: "最上", FirstName: "静香", EmojiName: "sizuka"},
	{LastName: "所", FirstName: "恵美", EmojiName: "megumi"},
	{FirstName: "ロコ", EmojiName: "roco"},
	{LastName: "天空橋", FirstName: "朋花", EmojiName: "tomoka"},
	{LastName: "北沢", FirstName: "志保", EmojiName: "shiho"},
	{LastName: "舞浜", FirstName: "歩", EmojiName: "ayumu"},
	{LastName: "二回", FirstName: "度千鶴", EmojiName: "chizuru"},
	{LastName: "真壁", FirstName: "瑞希", EmojiName: "mizuki"},
	{LastName: "百瀬", FirstName: "莉緒", EmojiName: "rio"},
	{LastName: "永吉", FirstName: "昴", EmojiName: "subaru"},
	{LastName: "周防", FirstName: "桃子", EmojiName: "momoko"},
	{FirstName: "ジュリア", EmojiName: "julia"},
	{LastName: "白石", FirstName: "紬", EmojiName: "tsumugi"},

	// Angel
	{LastName: "伊吹", FirstName: "翼", EmojiName: "tsubasa"},
	{LastName: "島原", FirstName: "エレナ", EmojiName: "elena"},
	{LastName: "箱崎", FirstName: "星梨花", EmojiName: "serika"},
	{LastName: "野々原", FirstName: "茜", EmojiName: "akane"},
	{LastName: "望月", FirstName: "杏奈", EmojiName: "anna"},
	{LastName: "木下", FirstName: "ひなた", EmojiName: "hinata"},
	{LastName: "馬場", FirstName: "このみ", EmojiName: "konomi"},
	{LastName: "大神", FirstName: "環", EmojiName: "tamaki"},
	{LastName: "豊川", FirstName: "風花", EmojiName: "fuka"},
	{LastName: "宮尾", FirstName: "美也", EmojiName: "miya"},
	{LastName: "篠宮", FirstName: "可憐", EmojiName: "karen"},
	{LastName: "北巻", FirstName: "麗花", EmojiName: "reika"},
	{LastName: "桜守", FirstName: "歌織", EmojiName: "kaori"},

	// Guests
	{FirstName: "詩花", EmojiName: "shika"},

	// Staffs
	{LastName: "音無", FirstName: "小鳥", EmojiName: "kotori"},
	{LastName: "青羽", FirstName: "美咲", EmojiName: "misaki"},
}

func chanAndSan(name string) string {
	return name + "ちゃん|" + name + "さん"
}

func emojinize(name string) string {
	return EmojiPrefix + name + EmojiSuffix
}

func generateRegexp(idols *[]Idol) {
	for i := range *idols {
		var err error
		last := (*idols)[i].LastName
		first := (*idols)[i].FirstName

		nameRule := chanAndSan(first) + "|" + last + first + "「"
		if (*idols)[i].LastName == "" {
			(*idols)[i].Regex, err = regexp.Compile("(" + nameRule + ")")
		} else {
			(*idols)[i].Regex, err = regexp.Compile("(" + nameRule + "|" + chanAndSan(last) + ")")
		}

		if err != nil {
			panic(err)
		}
	}
}

func canInsert(text, emoji string, index int) bool {
	if index == 0 {
		return true
	}

	if emoji == emojinize("ayumu") && string(text[index-len("雪"):index]) == "雪" {
		return false
	}

	if string(text[index-1]) == "S" {
		return false
	}

	return true
}

func insert(text, emoji string, index int) string {
	if len(text) < index {
		return text
	}

	if !canInsert(text, emoji, index) {
		return text
	}

	if string(text[index]) == " " { // if space
		text = text[:index] + emoji + text[index:] // insert
	}

	if text[index] == 10 { // if newline
		emoji = "\n" + emoji                         // add newline
		text = text[:index] + emoji + text[index+1:] // ignores old newline
	} else {
		text = text[:index] + emoji + text[index:] // insert
	}

	return text
}

func insertSR(text string) string {
	sr, err := regexp.Compile(`SR`)
	if err != nil {
		alertToOwner("REGEXP COMPILE ERROR")
		return text
	}

	return insertAll(text, "gasha_yellow", sr)
}

func insertAll(text, emojiName string, re *regexp.Regexp) string {
	indexes := re.FindAllStringIndex(text, -1) // find
	if indexes == nil {
		return text
	}

	emoji := emojinize(emojiName) // :mltd_name:
	addedLen := 0
	for _, index := range indexes {
		index[0] += addedLen
		textInserted := insert(text, emoji, index[0])
		if len(textInserted) == len(text) {
			continue
		}

		addedLen += len(textInserted) - len(text)
		text = textInserted
	}

	return text
}

func insertEmoji(text string) string {
	for i := range idolTable {
		text = insertAll(text, idolTable[i].EmojiName, idolTable[i].Regex)
	}

	text = insertSR(text)
	r := strings.NewReplacer(EasyIcons...) // types replacer
	text = r.Replace(text)                 // insert types emoji

	return text
}
