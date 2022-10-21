package bot

import (
	"log"
	"strings"
	"time"
	"topmusicstreaming/utils"
	"unicode/utf8" //nolint

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func Tweet(body string) {
	config := oauth1.NewConfig("***", "***")
	token := oauth1.NewToken("***", "***")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	tweet, resp, err := client.Statuses.Update(body, nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}

func TweetPosition(data int) string {
	if data == 1 {
		return "1️⃣"
	} else if data == 2 {
		return "2️⃣"
	} else if data == 3 {
		return "3️⃣"
	} else if data == 4 {
		return "4️⃣"
	}
	return "5️⃣"
}

func TweetEvolution(data string) string {
	if data == "+" {
		return "⬆️"
	} else if data == "=" {
		return "➡️"
	} else if data == "-" {
		return "⬇️"
	}
	return "🆕"
}

func TweetHeader(data string) string {
	paris, _ := time.LoadLocation("Europe/Paris")
	dt := time.Now().In(paris)
	if data == "us" {
		return "🇺🇸 TOP UNITED STATE - " + dt.Format("01/02/2006")
	} else if data == "fr" {
		return "🇫🇷 TOP FRANCE - " + dt.Format("01/02/2006")
	} else if data == "de" {
		return "🇩🇪 TOP GERMANY - " + dt.Format("01/02/2006")
	} else if data == "es" {
		return "🇪🇸 TOP SPAIN - " + dt.Format("01/02/2006")
	} else if data == "pt" {
		return "🇵🇹 TOP PORTUGAL - " + dt.Format("01/02/2006")
	} else if data == "it" {
		return "🇮🇹 TOP ITALY - " + dt.Format("01/02/2006")
	}
	return "🌎 TOP WORLD - " + dt.Format("01/02/2006")
}

func TweetHashtag(name1 string, name2 string, name3 string, name4 string, name5 string, runes int) string {
	hashtag := ""
	countRunes := utf8.RuneCountInString("#"+utils.TrimTweet(name1)) + runes
	if countRunes <= 280 {
		hashtag += "#" + utils.TrimTweet(name1)
	}
	if strings.ToLower(name2) != strings.ToLower(name1) {
		countRunes += utf8.RuneCountInString(" #" + utils.TrimTweet(name2))
		if countRunes <= 280 {
			hashtag += " #" + utils.TrimTweet(name2)
		}
	}
	if strings.ToLower(name3) != strings.ToLower(name1) && strings.ToLower(name3) != strings.ToLower(name2) {
		countRunes += utf8.RuneCountInString(" #" + utils.TrimTweet(name3))
		if countRunes <= 280 {
			hashtag += " #" + utils.TrimTweet(name3)
		}
	}
	if strings.ToLower(name4) != strings.ToLower(name1) && strings.ToLower(name4) != strings.ToLower(name2) && strings.ToLower(name4) != strings.ToLower(name3) {
		countRunes += utf8.RuneCountInString(" #" + utils.TrimTweet(name4))
		if countRunes <= 280 {
			hashtag += " #" + utils.TrimTweet(name4)
		}
	}
	if strings.ToLower(name5) != strings.ToLower(name1) && strings.ToLower(name5) != strings.ToLower(name2) && strings.ToLower(name5) != strings.ToLower(name3) && strings.ToLower(name5) != strings.ToLower(name4) {
		countRunes += utf8.RuneCountInString(" #" + utils.TrimTweet(name5))
		if countRunes <= 280 {
			hashtag += " #" + utils.TrimTweet(name5)
		}
	}
	return hashtag
}
