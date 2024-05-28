package telegram

type RequestEditCaption struct {
	ChatID      int         `json:"chat_id"`
	MessageID   int         `json:"message_id"`
	Caption     string      `json:"caption"`
	ParseMode   ParseMode   `json:"parse_mode"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type RequestEditMedia struct {
	ChatID      int         `json:"chat_id"`
	MessageID   int         `json:"message_id"`
	Media       InputMedia  `json:"media"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type InputMedia struct {
	Type        string      `json:"type"`
	Media       string      `json:"media"`
	Caption     string      `json:"caption"`
	ParseMode   ParseMode   `json:"parse_mode"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}
