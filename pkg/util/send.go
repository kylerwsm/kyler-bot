package util

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// SendToUser sends response to user.
func SendToUser(chatID int, text string) error {
	if token, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN"); ok {
		_, err := http.PostForm(
			fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
			url.Values{
				"chat_id": {strconv.Itoa(chatID)},
				"text":    {text},
			})
		return err
	}
	return errors.New("TELEGRAM_TOKEN environment variable is not defined")
}
