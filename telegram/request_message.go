package telegram

type RequestMessage struct {
	ChatID              int             `json:"chat_id"`
	Text                string          `json:"text"`
	Entities            []MessageEntity `json:"entities"`
	ReplyParametes      ReplyParametes  `json:"reply_parametes"`
	ReplyMarkup         *ReplyMarkup    `json:"reply_markup"`
	ParseMode           ParseMode       `json:"parse_mode"`
	DisableNotification bool            `json:"disable_notification"`
	ProtectContent      bool            `json:"protect_content"`
}

type ParseMode string

const (
	MarkdownV2 ParseMode = "MarkdownV2"
	HTML       ParseMode = "HTML"
)

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboard `json:"inline_keyboard"`
}
type InlineKeyboard struct {
	Text                         string `json:"text"`
	URL                          string `json:"url"`
	CallbackData                 string `json:"callback_data"`
	SwitchInlineQuery            string `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`
	Pay                          bool   `json:"pay"  bson:"pay"`
}

type ReplyParametes struct {
	MessageID int `json:"message_id"`
	ChatID    int `json:"chat_id"`
}
