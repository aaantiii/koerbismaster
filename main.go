package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

const PROD = true

func init() {
	log.SetPrefix("[BOT] ")
	log.SetFlags(log.Ldate | log.Ltime)

	if err := initEnv(); err != nil {
		log.Fatalf("Failed to init environment variables: %v", err)
	}
}

func main() {
	session, err := newDiscordSession()
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

func newDiscordSession() (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + DISCORD_CLIENT_SECRET.Value())
	if err != nil {
		return nil, err
	}

	if err = session.Open(); err != nil {
		return nil, err
	}

	session.AddHandler(handleMessageCreate)
	return session, nil
}
