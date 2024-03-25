package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/aaantiii/koerbismaster"
)

func init() {
	slog.SetDefault(koerbismaster.NewLogger())
	if err := koerbismaster.LoadEnv(); err != nil {
		slog.Error("Failed to init environment variables", slog.Any("err", err))
		os.Exit(1)
	}
}

func main() {
	session, err := koerbismaster.NewClient()
	if err != nil {
		slog.Error("Failed to create discord session", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Info("Bot is logged in and running. Press CTRL-C to exit.", slog.String("username", session.State.User.Username))

	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, os.Interrupt)
	<-shutdownSig

	slog.Info("Gracefully shutting down...")
	if err = session.Close(); err != nil {
		slog.Error("Failed to close discord session", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Info("Shutdown successful.")
}
