package webhook

import (
	"context"
	"errors"
	"fmt"

	"github.com/seosoojin/go-telegram/telegram"
)

var (
	ErrNotFound = errors.New("command not found")
)

func ErrorHandler(ctx context.Context, client telegram.Client, event *telegram.Event, err error) {
	id := 0

	switch {
	case event.CallbackQuery != nil:
		id = event.CallbackQuery.Message.Chat.ID
	case event.Message != nil:
		id = event.Message.Chat.ID
	default:
		return
	}

	if id == 0 {
		return
	}

	client.SendMessage(telegram.RequestMessage{
		ChatID: id,
		Text:   fmt.Sprintf("An error occurred: %v", err),
	})
}
