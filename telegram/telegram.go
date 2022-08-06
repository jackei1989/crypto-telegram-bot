package telegram

import (
	"log"
	"net/http"
	"net/url"
)

func SendMessage(chatID string, text string) {

	token := "5490111796:AAHfwevpxv7oRVuQmLk-RDhPw491bLbmnIE"

	telegramUrl, err := url.Parse("https://api.telegram.org/bot" + token + "/sendMessage")

	value := telegramUrl.Query()
	value.Add("chat_id", chatID)
	value.Add("text", text)
	value.Add("parse_mode", "html")
	telegramUrl.RawQuery = value.Encode()
	if err != nil {
		log.Println(err)
	}

	response, err := http.Get(telegramUrl.String())
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
}
