package telegram

type Message struct {
	MessageID       int64           `json:"message_id"`
	From            From            `json:"from"`
	Chat            Chat            `json:"chat"`
	Date            int64           `json:"date"`
	Text            string          `json:"text"`
	Photo           []Photo         `json:"photo"`
	Caption         string          `json:"caption"`
	CaptionEntities []CaptionEntity `json:"caption_entities"`
	ReplyMarkup     ReplyMarkup     `json:"reply_markup"`
}

type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID int `json:"id"`
}
