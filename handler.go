package koerbismaster

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	countdown   = time.Second * 15
	deleteAfter = time.Second * 10
)

func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Failed to handle MessageCreate: %v", err)
			}
		}()
		handle(s, m)
	}()
}

func checkMessageEmbed(m *discordgo.Message) bool {
	if len(m.Embeds) == 0 {
		return false
	}

	return m.Embeds[0].Title == DISCORD_EVENET_SYS_CONTENT.Value()
}

// for debugging
func checkMessageContent(m *discordgo.Message) bool {
	isMatch := m.Content == DISCORD_EVENET_SYS_CONTENT.Value()
	if !isMatch {
		log.Println("Content does not match.")
	}

	return isMatch
}

func handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != DISCORD_EVENT_SYS_CLIENT_ID.Value() || m.ChannelID != DISCORD_EVENT_SYS_CHANNEL_ID.Value() {
		return
	}

	if !checkMessageEmbed(m.Message) {
		return
	}

	log.Printf("Detected event system action within %s.", time.Since(m.Timestamp).Round(time.Millisecond))
	res, err := sendPingMessage(s, m.Message)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return
	}

	time.Sleep(deleteAfter)
	if err = s.ChannelMessageDelete(DISCORD_PING_CHANNEL_ID.Value(), res.ID); err != nil {
		log.Printf("Failed to delete ping message: %v", err)
	}
}
