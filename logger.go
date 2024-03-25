package koerbismaster

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: getLogLevel(),
	}))
}

func getLogLevel() slog.Level {
	switch MODE.Value() {
	case "PROD":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	default:
		return slog.LevelDebug
	}
}
