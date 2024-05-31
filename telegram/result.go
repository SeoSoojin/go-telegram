package telegram

type Result struct {
	OK     bool    `json:"ok"`
	Result Message `json:"result"`
}
