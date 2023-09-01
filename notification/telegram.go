package notification

import (
	"bus_listener/env"
	"fmt"
	"io"
	"net/http"
)

type Telegram struct {
	url    string
	chatId string
}

func NewTelegram() *Telegram {
	token := env.Telegram.Token
	chatId := env.Telegram.UserId
	return &Telegram{
		url:    fmt.Sprintf("https://api.telegram.org/bot%s", token),
		chatId: chatId,
	}
}

func (t Telegram) getUrl(link string) string {
	return fmt.Sprintf("%s/%s", t.url, link)
}

func (t Telegram) Send(message Message) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, t.getUrl("sendMessage"), nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("text", message.Text)
	q.Add("chat_id", t.chatId)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return err
	}

	fmt.Println(req.URL.String())

	defer func() {
		_ = resp.Body.Close()
	}()
	text, _ := io.ReadAll(resp.Body)
	textStr := string(text)
	_ = textStr
	return nil
}
