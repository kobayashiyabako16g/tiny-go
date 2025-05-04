package log

import (
	"log/slog"
	"os"
	"strings"
)

var Logger *slog.Logger

func init() {
	level := getLogLevel(os.Getenv("LOG_LEVEL"))
	format := strings.ToLower(os.Getenv("LOG_FORMAT"))

	var handler slog.Handler
	switch format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}

	Logger = slog.New(handler)
}

func getLogLevel(envLevel string) slog.Level {
	switch strings.ToLower(envLevel) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo

	}
}
