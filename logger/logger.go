package logger

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Level  string // "debug", "info", "warn", "error"
	Format string // "json" or "text"
}

// NewLogger creates a new structured logger
func NewLogger(config *Config) *slog.Logger {
	level := parseLevel(config.Level)

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: level,
	}

	if config.Format == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

// parseLevel converts string level to slog.Level
func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
