package koerbismaster

import (
	"fmt"
	"log"
	"log/slog"
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
		if _, err = s.ChannelMessageEdit(msg.ChannelID, msg.ID, createPingMessage(i)); err != nil {
			log.Printf("Failed to edit message to update countdown: %v", err)
			continue
		}
		time.Sleep(time.Second - time.Since(now))
	}

	return err
}

func sendMessageLinks(s *discordgo.Session, m *discordgo.Message) {
	msg := ""
	url := fmt.Sprintf("https://discord.com/channels/%s/%s/%s\n", m.GuildID, m.ChannelID, m.ID)
	for i := 0; i < 3; i++ {
		msg += url
	}

	res, err := s.ChannelMessageSend(DISCORD_PING_CHANNEL_ID.Value(), msg)
	if err != nil {
		slog.Error("Failed to send message with links to event message.", slog.Any("err", err))
	}
	go deleteMessageAfter(s, res, deleteAfter)

	slog.Info("Sent message with links to event message.")
}

func createPingMessage(secondsLeft int) string {
	msg := fmt.Sprintf("%s\nDer Link zur Nachricht wird gleich gesendet...\n", mentionRole(DISCORD_PING_ROLE_ID.Value()))
	if secondsLeft > 0 {
		msg += fmt.Sprintf("Noch **%d Sekunden** verfügbar.", secondsLeft)
	} else {
		msg += "**Event vorbei**"
	}
	return msg
}

func deleteMessageAfter(s *discordgo.Session, m *discordgo.Message, d time.Duration) {
	time.Sleep(d)
	if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
		slog.Error("Failed to delete message.", slog.Any("err", err))
	}
}

func mentionRole(roleID string) string {
	return fmt.Sprintf("<@&%s>", roleID)
}
