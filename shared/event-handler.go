package shared

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
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
	writer := os.Stdout

	// create a new logger
	logger := slog.New(tint.NewHandler(writer, nil))

	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(writer, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	return &EventHandler{
		logger: logger,
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
