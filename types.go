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
	Text 		              string     `json:"text"` 
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
	Description         string     `json:"description,omitempty"`
	InviteLink          string     `json:"invite_link,omitempty"`
}


type Update struct {
	ID int64 `json:"update_id"`
	Payload *Message `json:"message"`
}


func (c Chat) IsPrivate() bool {
	return c.Type == "private"
}
func (c Chat) IsGroup() bool {
	return c.Type == "group"
}

func (c Chat) IsSuperGroup() bool {
	return c.Type == "supergroup"
}

func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}


func (m *Message) IsCommand() bool {
	return m.Text != "" && strings.HasPrefix(m.Text,"/")
}


func (m *Message) Command() string {
	if !m.IsCommand() {
		return ""
	}

	command := strings.SplitN(m.Text, " ", 2)[0][1:]

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	split := strings.SplitN(m.Text, " ", 2)
	if len(split) != 2 {
		return ""
	}

	return split[1]
}
