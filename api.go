package telebot

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
	"errors"
)

const (
	BASE_URL = "https:/api.telegram.org/bot%s/%s"
)

func (b *Bot) getMe() (User, error) {
	url := fmt.Sprintf(BASE_URL, b.Token, "getMe")

	resp, err := getContent(url)

	if err != nil {
		return User{}, err
	}

	var botInfo struct {
		Ok bool 
		Result User 
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

func (b *Bot) SendToChannel(channelName string, text string)(error) {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", b.Token, channelName, text)

	_, err := getContent(url)

	if err != nil {
		return err 
	}

	return nil 
}



func (b *Bot) GetUpdates(offset int64, timeout int64)([]Update,error) {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d&timeout=%d", b.Token, offset, timeout)
	
	resp, err := http.Get(url)

	var updatesReceived struct {
		Ok bool 
		Result []Update 
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

func getContent(url string)([]byte, error) {
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

func (b *Bot) SetWebhook(config WebhookConfig) bool {
	// a https url to send updates to
	// optional: ssl certificate 
	// optional : max connections 
	// allowed updates: optional

	if config.Certificate == nil {
		v := url.Values{}
		v.Add("url", config.URL.String())
	}


	return true
}

func (b *Bot) DeleteWebhook() bool {
	return true 
}