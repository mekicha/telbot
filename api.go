package telebot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (b *Bot) getMe() (User, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", b.Token)

	resp, err := getContent(url)

	if err != nil {
		return User{}, err
	}

	var botInfo struct {
		Ok          bool
		Result      User
		Description string
	}

	err = json.Unmarshal(resp, &botInfo)
	if err != nil {
		return User{}, err
	}

	if !botInfo.Ok {
		return User{}, errors.New("bad response")
	}

	return botInfo.Result, nil

}

// SendMessage sends message to a user
func (b *Bot) SendMessage(chatID int64, text string) (Message, error) {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", b.Token, chatID, text)

	resp, err := getContent(url)

	if err != nil {
		return Message{}, err
	}

	var m Message
	err = json.Unmarshal(resp, &m)

	if err != nil {
		return Message{}, err
	}

	return m, nil

}

// SendToChannel sends message to a registered channel @channelName
func (b *Bot) SendToChannel(channelName string, text string) error {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", b.Token, channelName, text)

	_, err := getContent(url)

	if err != nil {
		return err
	}

	return nil
}

// GetUpdates gets update from telegram via long polling
func (b *Bot) GetUpdates(offset int64, timeout int64) ([]Update, error) {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d&timeout=%d", b.Token, offset, timeout)

	resp, err := http.Get(url)

	var updatesReceived struct {
		Ok          bool
		Result      []Update
		Description string
	}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&updatesReceived)

	if err != nil {
		return updatesReceived.Result, err
	}

	if !updatesReceived.Ok {
		return updatesReceived.Result, errors.New(updatesReceived.Description)
	}

	return updatesReceived.Result, nil
}

func getContent(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// SetWebhook sets a url as webhook. This disables GetUpdates
func (b *Bot) SetWebhook(url string) bool {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", b.Token, url)
	resp, err := http.Get(endpoint)
	if err != nil {
		return false
	}
	var res APIResponse

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&res)

	if !res.Ok {
		return false
	}
	return true
}

// DeleteWebhook removes webhook url
func (b *Bot) DeleteWebhook() bool {
	url := fmt.Sprintf("https://api.telegram.com/bot%s/deleteWebhook", b.Token)
	_, err := http.Get(url)
	if err != nil {
		return false
	}
	return true
}

// ListenForWebhook gets updates sent to the webhook url
func (b *Bot) ListenForWebhook(pattern string) UpdatesChannel {
	ch := make(chan Update, 10)

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadAll(r.Body)

		var update Update
		json.Unmarshal(bytes, &update)

		ch <- update
	})

	return ch
}
