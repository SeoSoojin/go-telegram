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
	client.SendMessage(telegram.RequestMessage{
		ChatID: event.Message.Chat.ID,
		Text:   fmt.Sprintf("An error occurred: %v", err),
	})
}
