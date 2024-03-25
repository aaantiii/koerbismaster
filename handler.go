package koerbismaster

import (
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	countdown   = time.Second * 15
	deleteAfter = time.Second * 30
)

func handleTypingStart(s *discordgo.Session, t *discordgo.TypingStart) {
	defer handleRecovery("TypingStart")
	if t.UserID != DISCORD_EVENT_SYS_CLIENT_ID.Value() || t.ChannelID != DISCORD_EVENT_SYS_CHANNEL_ID.Value() {
		return
	}

	if err := sendPingMessage(s, time.Unix(int64(t.Timestamp), 0)); err != nil {
		slog.Error("Failed to send ping message.", slog.Any("err", err), slog.String("handler", "TypingStart"))
		return
	}
}

func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer handleRecovery("MessageCreate")
	if m.ChannelID != DISCORD_EVENT_SYS_CHANNEL_ID.Value() {
		return
	}
	if m.Author.ID != DISCORD_EVENT_SYS_CLIENT_ID.Value() {
		stats.IncMessages()
		return
	}
	if !checkMessage(m.Message) {
		slog.Debug("Message does not match.", slog.String("handler", "MessageCreate"))
		return
	}
	slog.Info("Detected event from EventSystem.", slog.String("handler", "MessageCreate"))

	go sendMessageLinks(s, m.Message)
	if err := stats.Save(); err != nil {
		slog.Error("Failed to save stats.", slog.Any("err", err), slog.String("handler", "MessageCreate"))
	}
	stats.Reset()
	slog.Info("Saved stats.", slog.String("handler", "MessageCreate"))
}

func handleRecovery(handlerName string) {
	if err := recover(); err != nil {
		slog.Error("Failed to recover event.", slog.String("handler", handlerName), slog.Any("err", err))
	}
}

func checkMessage(m *discordgo.Message) bool {
	if MODE.Value() == "PROD" {
		return checkMessageEmbed(m)
	}
	return checkMessageContent(m)
}

func checkMessageEmbed(m *discordgo.Message) bool {
	if len(m.Embeds) == 0 {
		return false
	}

	return m.Embeds[0].Title == DISCORD_EVENT_SYS_CONTENT.Value()
}

// for debugging
func checkMessageContent(m *discordgo.Message) bool {
	return m.Content == DISCORD_EVENT_SYS_CONTENT.Value()
}
