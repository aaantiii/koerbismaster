package koerbismaster

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	countdown   = time.Second * 16
	deleteAfter = time.Second * 30
)

func handleTypingStart(s *discordgo.Session, m *discordgo.TypingStart) {
	go func() {
		defer handleRecovery("TypingStart")
		if m.UserID != DISCORD_EVENT_SYS_CLIENT_ID.Value() || m.ChannelID != DISCORD_EVENT_SYS_CHANNEL_ID.Value() {
			return
		}

		if err := sendPingMessage(s, time.Unix(int64(m.Timestamp), 0)); err != nil {
			log.Printf("Failed to send ping message: %v", err)
			return
		}
	}()
}

func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	go func() {
		defer handleRecovery("MessageCreate")
		if m.Author.ID != DISCORD_EVENT_SYS_CLIENT_ID.Value() || m.ChannelID != DISCORD_EVENT_SYS_CHANNEL_ID.Value() {
			return
		}

		if PROD {
			if !checkMessageEmbed(m.Message) {
				return
			}
		} else {
			if !checkMessageContent(m.Message) {
				return
			}
		}

		log.Println("MessageCreate: EventSystem sent message.")

		if err := sendMessageLinks(s, m.Message); err != nil {
			log.Printf("MessageCreate: failed to send message: %v", err)
			return
		}
		log.Printf("MessageCreate: sent message with links.")
	}()
}

func handleRecovery(handlerName string) {
	if err := recover(); err != nil {
		log.Printf("Failed to handle %s: %v", handlerName, err)
	}
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
