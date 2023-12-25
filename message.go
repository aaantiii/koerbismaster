package koerbismaster

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func sendPingMessage(s *discordgo.Session, received time.Time) error {
	msg, err := s.ChannelMessageSend(DISCORD_PING_CHANNEL_ID.Value(), mentionRole(DISCORD_PING_ROLE_ID.Value()))
	if err != nil {
		return err
	}
	log.Printf("TypingStart: sent ping within %s.", time.Since(received.Round(time.Millisecond)))
	go deleteMessageAfter(s, msg, deleteAfter)

	for i := int(countdown.Seconds()) - 1; i >= 0; i-- {
		now := time.Now()
		if _, err = s.ChannelMessageEdit(msg.ChannelID, msg.ID, pingMessage(i)); err != nil {
			log.Printf("Failed to edit message to update countdown: %v", err)
			continue
		}
		time.Sleep(time.Second - time.Since(now))
	}

	return err
}

func sendMessageLinks(s *discordgo.Session, m *discordgo.Message) error {
	msg := ""
	url := fmt.Sprintf("https://discord.com/channels/%s/%s/%s\n", m.GuildID, m.ChannelID, m.ID)
	for i := 0; i < 3; i++ {
		msg += url
	}

	res, err := s.ChannelMessageSend(DISCORD_PING_CHANNEL_ID.Value(), msg)
	if err != nil {
		return err
	}
	go deleteMessageAfter(s, res, deleteAfter)

	return nil
}

func pingMessage(secondsLeft int) string {
	msg := fmt.Sprintf("%s\nDer Link zur Nachricht wird gleich gesendet...\n", mentionRole(DISCORD_PING_ROLE_ID.Value()))
	if secondsLeft > 0 {
		msg += fmt.Sprintf("Noch **%d Sekunden** verfügbar.", secondsLeft)
	} else {
		msg += "**Event beendet.**"
	}

	return msg
}

func deleteMessageAfter(s *discordgo.Session, m *discordgo.Message, d time.Duration) {
	time.Sleep(d)
	if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
		log.Printf("Failed to delete message: %v", err)
	}
}

func mentionRole(role string) string {
	return fmt.Sprintf("<@&%s>", role)
}
