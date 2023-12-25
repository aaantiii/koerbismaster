package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/aaantiii/koerbismaster"
)

func init() {
	log.SetPrefix("[BOT] ")
	log.SetFlags(log.Ldate | log.Ltime)

	if err := koerbismaster.LoadEnv(); err != nil {
		log.Fatalf("Failed to init environment variables: %v", err)
	}
}

func main() {
	session, err := koerbismaster.NewClient()
	if err != nil {
		log.Fatalf("Failed to create discord session: %v", err)
	}
	log.Printf("Bot is logged in as %s and running. Press CTRL-C to exit.", session.State.User.Username)

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt)
	<-shutdownSignal

	log.Println("Gracefully shutting down...")
	if err = session.Close(); err != nil {
		log.Fatalf("Failed to close discord session: %v", err)
	}
}
