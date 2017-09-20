package telebot 

import (
	"strings"
)


type Message struct {
	MessageID             int        `json:"message_id"`
	From                  *User      `json:"from"` 
	Date                  int        `json:"date"`
	Chat                  *Chat      `json:"chat"`
	ForwardFrom           *User      `json:"forward_from"`      
	ForwardFromChat       *Chat      `json:"forward_from_chat"` 
	ForwardFromMessageID  int        `json:"forward_from_message_id"`
	ForwardDate           int        `json:"forward_date"` 
	ReplyToMessage        *Message   `json:"reply_to_message"`
	EditDate              int        `json:"edit_date"`
	Text string                      `json:"text"` 
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	IsBot bool `json:"is_bot"`
	Language string `json:"language_code"`
}


type Chat struct {
	ID                  int64      `json:"id"`
	Type                string     `json:"type"`
	Title               string     `json:"title"`                         
	UserName            string     `json:"username"`                      
	FirstName           string     `json:"first_name"`                    
	LastName            string     `json:"last_name"`                     
	AllMembersAreAdmins bool       `json:"all_members_are_administrators"`
	Photo               *ChatPhoto `json:"photo"`
	Description         string     `json:"description,omitempty"`
	InviteLink          string     `json:"invite_link,omitempty"`
}

func (c Chat) IsPrivate() bool {
	return c.Type == "private"
}
unc (c Chat) IsGroup() bool {
	return c.Type == "group"
}

func (c Chat) IsSuperGroup() bool {
	return c.Type == "supergroup"
}

func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}


func (m *Message) IsCommand() bool {
	return m.Text != "" && strings.HasPrefix(m.Text,'/')
}