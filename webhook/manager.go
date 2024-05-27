package webhook

import (
	"context"

	"github.com/seosoojin/go-telegram/telegram"
)

type Command func(ctx context.Context, client *telegram.Client, event *telegram.Event, args ...string) error

type Manager struct {
	commands map[string]Command
}

func NewManager() *Manager {
	return &Manager{
		commands: make(map[string]Command),
	}
}

func (h *Manager) AddCommand(name string, command Command) {
	h.commands[name] = command
}

func (h *Manager) GetCommand(name string) Command {
	return h.commands[name]
}
