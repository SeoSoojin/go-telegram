package telegram

type RequestPhoto struct {
	ChatID         int            `json:"chat_id"`
	Photo          string         `json:"photo"`
	Caption        string         `json:"caption"`
	ReplyParametes ReplyParametes `json:"reply_parametes"`
	ReplyMarkup    *ReplyMarkup   `json:"reply_markup"`
	ParseMode      ParseMode      `json:"parse_mode"`
}

type RequestAnimation struct {
	ChatID         int            `json:"chat_id"`
	Animation      string         `json:"animation"`
	Caption        string         `json:"caption"`
	ReplyParametes ReplyParametes `json:"reply_parametes"`
	ReplyMarkup    *ReplyMarkup   `json:"reply_markup"`
	ParseMode      ParseMode      `json:"parse_mode"`
}
