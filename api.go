package telebot

func (b *Bot) getMe() (User, error) {
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

func (b *Bot) sendMessage(chatID int64, text string) (Message, error) {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", b.Token, chatID, text)

	var buf bytes.Buffer
	
		resp, err := http.Post(url, "application/json", &buf)

		if err != nil {
			return Message{}, err 
		}

		var message Message

		json.Unmarshal(resp.Body, &message)

		return message, nil 

}