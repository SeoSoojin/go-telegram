package webhook

import (
	"io"
	"net/http"
	"strings"

	"github.com/seosoojin/go-telegram/telegram"
)

type Server interface {
	Route(name string, command Command)
	Run(port string) error
}

type server struct {
	tgClient telegram.Client
	manager  *Manager
}

func NewServer(tgClient telegram.Client) Server {
	return &server{
		tgClient: tgClient,
		manager:  NewManager(),
	}
}

func (s *server) Route(name string, command Command) {
	s.manager.AddCommand(name, command)
}

func (s *server) Run(port string) error {
	http.HandleFunc("/webhook", s.handler)
	return http.ListenAndServe(":"+port, nil)
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {

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

	command := s.manager.GetCommand(commandName)
	if command == nil {
		http.Error(w, "command not found", http.StatusNotFound)
		return
	}

	err = command(r.Context(), s.tgClient, event, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
