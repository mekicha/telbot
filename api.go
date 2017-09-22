package telebot

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
)


func (b *Bot) getMe() (User, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.Token, "getMe")


	resp, err := http.Get(url)

	if err != nil {
		return User{}, err 
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var user User 
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, nil 
	}
	return user, nil

}

func (b *Bot) sendMessage(chatID int64, text string) (Message, error) {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", b.Token, chatID, text)

	resp, err := http.Get(url)
	
		if err != nil {
			return Message{}, err 
		}
	
		defer resp.Body.Close()
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Message{}, err
		}
	
		var m Message
		err = json.Unmarshal(body, &m)

		if err != nil {
			return Message{}, nil 
		}

		return m, nil

}



func (b *Bot) getUpdates(offset int64, timeout int64)([]Update,error) {

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