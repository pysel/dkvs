package shared

import (
	"log/slog"
)

// EventHandler is a struct that handles events by logging them.
type EventHandler struct {
	logger *slog.Logger
}

type Event interface {
	Severity() string
	Message() string
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		logger: slog.Default(),
	}
}

func (e *EventHandler) Handle(event Event) {
	switch event.Severity() {
	case "info":
		e.logger.Info(event.Message())
	case "warning":
		e.logger.Warn(event.Message())
	case "error":
		e.logger.Error(event.Message())
		// Maybe panic here?
	}
}
