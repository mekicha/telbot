package telebot

import (
	"io/ioutil"
	"bytes"
	"net/http"
)

type Bot struct {
	Token string 
	Owner User
}

func NewBot(token string) (*Bot, error){
	bot := &Bot{
				 Token: token
	}
	
	owner, err := bot.getMe()
	if err != nil {
		return nil, err 
	}
	bot.Owner = owner 
	return bot, nil 
}


func (*b Bot) getMe() (User, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.Token, "getMe")

	var buf bytes.Buffer
	
	resp, err := http.Post(url, "application/json", &buf)
	if err != nil {
		return []byte{}, errors.Wrap(err, "http.Post failed")
	}
	resp.Close = true
	defer resp.Body.Close()
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, wrapSystem(err)
	}

return json, nil
}
