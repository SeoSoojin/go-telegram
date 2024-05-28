package webhook

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/seosoojin/go-telegram/telegram"
)

type Server interface {
	Route(name string, command Command, opts ...RouteOptions)
	Run(port string) error
}

type server struct {
	tgClient        telegram.Client
	manager         *Manager
	callbackManager *Manager
}

func NewServer(tgClient telegram.Client) Server {
	return &server{
		tgClient:        tgClient,
		manager:         NewManager(),
		callbackManager: NewManager(),
	}
}

func (s *server) Route(name string, command Command, opts ...RouteOptions) {
	routeOpts := newRouteOptions(opts...)

	if routeOpts.callbacks != nil {
		for name, command := range routeOpts.callbacks {
			s.callbackManager.AddCommand(name, command)
		}
	}

	s.manager.AddCommand(name, command)
}

func (s *server) Run(port string) error {
	http.HandleFunc("/webhook", s.handler)
	return http.ListenAndServe(":"+port, nil)
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	event, err := telegram.ParseEvent(b)
	if err != nil {
		ErrorHandler(r.Context(), s.tgClient, event, err)
		return
	}

	switch {
	case event.Message != nil:
		s.handleCommand(r.Context(), event, w)
	case event.CallbackQuery != nil:
		s.handleCallback(r.Context(), event, w)
	default:
		ErrorHandler(r.Context(), s.tgClient, event, ErrNotFound)
	}

}

func (s *server) handleCallback(ctx context.Context, event *telegram.Event, w http.ResponseWriter) {

	commandParts := strings.Split(event.Message.Text, " ")
	if len(commandParts) == 0 {
		ErrorHandler(ctx, s.tgClient, event, ErrNotFound)
		return
	}

	commandName := strings.Split(commandParts[0], "@")[0]
	args := commandParts[1:]

	command := s.callbackManager.GetCommand(commandName)
	if command == nil {
		ErrorHandler(ctx, s.tgClient, event, ErrNotFound)
		return
	}

	err := command(ctx, s.tgClient, event, args...)
	if err != nil {
		ErrorHandler(ctx, s.tgClient, event, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *server) handleCommand(ctx context.Context, event *telegram.Event, w http.ResponseWriter) {

	commandParts := strings.Split(event.Message.Text, " ")
	if len(commandParts) == 0 {
		ErrorHandler(ctx, s.tgClient, event, ErrNotFound)
		return
	}

	commandName := strings.Split(commandParts[0], "@")[0]
	args := commandParts[1:]

	command := s.manager.GetCommand(commandName)
	if command == nil {
		ErrorHandler(ctx, s.tgClient, event, ErrNotFound)
		return
	}

	err := command(ctx, s.tgClient, event, args...)
	if err != nil {
		ErrorHandler(ctx, s.tgClient, event, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
