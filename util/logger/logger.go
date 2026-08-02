package logger

import (
	"log/slog"
	"os"
)

// NewLogger return a logger. It take parameter isProd: is_production & addSrc: address_source.
func NewLogger(isProd, addSrc bool) *slog.Logger {

	level := slog.LevelDebug

	if isProd {
		level = slog.LevelInfo
		addSrc = false
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: addSrc,
	}

	var handler slog.Handler

	if isProd {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
