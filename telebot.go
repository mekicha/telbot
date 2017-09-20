package telebot

import (
	"io/ioutil"
	"bytes"
	"net/http"
	"fmt"
	"encoding/json"
)

type Bot struct {
	Token string 
	Owner User
}

func NewBot(token string) (*Bot, error){
	bot := &Bot{
	       Token: token,
	}
	
	owner, err := bot.getMe()
	if err != nil {
		return nil, err 
	}
	bot.Owner = owner 
	return bot, nil 
}

