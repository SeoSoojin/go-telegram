package webhook

import (
	"io"
	"net/http"
	"strings"

	"github.com/seosoojin/go-telegram/telegram"
)

type Client interface {
}

type client struct {
	tgClient *telegram.Client
	manager  *Manager
}

func NewClient(tgClient *telegram.Client) Client {
	return &client{
		tgClient: tgClient,
		manager:  NewManager(),
	}
}

func (c *client) Route(name string, command Command) {
	c.manager.AddCommand(name, command)
}

func (c *client) Run(port string) error {
	http.HandleFunc("/webhook", c.handler)
	return http.ListenAndServe(":"+port, nil)
}

func (c *client) handler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	event, err := telegram.ParseEvent(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commandParts := strings.Split(event.Message.Text, " ")
	if len(commandParts) == 0 {
		http.Error(w, "command not found", http.StatusNotFound)
		return
	}

	commandName := strings.Split(event.Message.Text, "@")[0]
	args := commandParts[1:]

	command := c.manager.GetCommand(commandName)
	if command == nil {
		http.Error(w, "command not found", http.StatusNotFound)
		return
	}

	err = command(r.Context(), c.tgClient, event, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
