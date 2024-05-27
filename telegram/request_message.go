package telegram

type RequestMessage struct {
	ChatID              int             `json:"chat_id"`
	Text                string          `json:"text"`
	Entities            []MessageEntity `json:"entities"`
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
