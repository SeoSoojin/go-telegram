package telegram

import "encoding/json"

type Event struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

func ParseEvent(data []byte) (*Event, error) {
	event := new(Event)
	if err := json.Unmarshal(data, event); err != nil {
		return nil, err
	}
	return event, nil
}
