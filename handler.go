package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Failed to handle message: %v", err)
			}
		}()
		handle(s, m)
	}()
}

func handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != DISCORD_EVENT_SYS_CLIENT_ID.Value() || m.ChannelID != DISCORD_EVENT_SYS_CHANNEL_ID.Value() {
		return
	}

	if len(m.Embeds) == 0 {
		return
	}

	embed := m.Embeds[0]
	if embed.Title != DISCORD_EVENET_SYS_CONTENT.Value() {
		return
	}

	log.Printf("Detected event system action within %s.", time.Since(m.Timestamp).Round(time.Millisecond))
	res, err := s.ChannelMessageSend(DISCORD_PING_CHANNEL_ID.Value(), pingMessage(m.GuildID, m.ChannelID, m.ID))
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
	log.Printf("Responded to event system action within %s.", time.Since(m.Timestamp).Round(time.Millisecond))

	time.Sleep(time.Second * 30)
	if err = s.ChannelMessageDelete(DISCORD_PING_CHANNEL_ID.Value(), res.ID); err != nil {
		log.Printf("Failed to delete message: %v", err)
	}
}

func handleTest(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != DISCORD_EVENT_SYS_CLIENT_ID.Value() || m.ChannelID != DISCORD_EVENT_SYS_CHANNEL_ID.Value() {
		log.Print("Author id or channel does not match.")
		return
	}

	if m.Content != DISCORD_EVENET_SYS_CONTENT.Value() {
		log.Printf("Content does not match: %s", m.Content)
		return
	}

	log.Printf("Detected event system action within %s.", time.Since(m.Timestamp).Round(time.Millisecond))
	res, err := s.ChannelMessageSend(DISCORD_PING_CHANNEL_ID.Value(), pingMessage(m.GuildID, m.ChannelID, m.ID))
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
	log.Printf("Responded to event system action within %s.", time.Since(m.Timestamp).Round(time.Millisecond))

	time.Sleep(time.Second * 5)
	if err = s.ChannelMessageDelete(DISCORD_PING_CHANNEL_ID.Value(), res.ID); err != nil {
		log.Printf("Failed to delete message: %v", err)
	}
}
