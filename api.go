package telebot

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
)

const (
	BASE_URL = "https://api.telegram.org/bot%s/%s"
)

func (b *Bot) getMe() (User, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", b.Token)

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

func (b *Bot) DeleteWebhook() bool {
	url := fmt.Sprintf(BASE_URL, b.Token, "deleteWebhook")
	_, err := http.Get(url)
	if err != nil {
		return false 
	}
	return true 
}

func (bot *Bot) ListenForWebhook(pattern string) UpdatesChannel {
	ch := make(chan Update, 10)

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadAll(r.Body)

		var update Update
		json.Unmarshal(bytes, &update)

		ch <- update
	})

	return ch
}