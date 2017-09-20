package telebot 


type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	IsBot bool `json:"is_bot"`
	Language string `json:"language_code"`
}