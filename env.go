package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVar string // EnvVar type represents an environment variable.

const (
	DISCORD_CLIENT_ID            EnvVar = "DISCORD_CLIENT_ID"
	DISCORD_CLIENT_SECRET        EnvVar = "DISCORD_CLIENT_SECRET"
	DISCORD_EVENT_SYS_CLIENT_ID  EnvVar = "DISCORD_EVENT_SYS_CLIENT_ID"
	DISCORD_EVENT_SYS_CHANNEL_ID EnvVar = "DISCORD_EVENT_SYS_CHANNEL_ID"
	DISCORD_EVENET_SYS_CONTENT   EnvVar = "DISCORD_EVENT_SYS_CONTENT"
	DISCORD_PING_CHANNEL_ID      EnvVar = "DISCORD_PING_CHANNEL_ID"
	DISCORD_PING_ROLE_ID         EnvVar = "DISCORD_PING_ROLE_ID"
)

// Value returns the value of the environment variable as string.
func (v EnvVar) Value() string {
	return os.Getenv(v.Name())
}

// Name returns the name of the environment variable.
func (v EnvVar) Name() string {
	return string(v)
}

func initEnv() error {
	filename := ".env.development"
	if PROD {
		filename = ".env.production"
	}

	if err := godotenv.Load(filename); err != nil {
		return err
	}

	required := []EnvVar{
		DISCORD_CLIENT_ID,
		DISCORD_CLIENT_SECRET,
		DISCORD_EVENT_SYS_CLIENT_ID,
		DISCORD_EVENT_SYS_CHANNEL_ID,
		DISCORD_EVENET_SYS_CONTENT,
		DISCORD_PING_CHANNEL_ID,
		DISCORD_PING_ROLE_ID,
	}

	for _, v := range required {
		if _, found := os.LookupEnv(v.Name()); !found {
			return fmt.Errorf("required env variable '%s' not set", v.Name())
		}
	}

	log.Println("All required env variables are set.")
	return nil
}
