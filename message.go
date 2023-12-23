package koerbismaster

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func pingMessage(m *discordgo.Message, secondsLeft int) string {
	msg := fmt.Sprintf("<@&%s>\n", DISCORD_PING_ROLE_ID.Value())
	url := fmt.Sprintf("https://discord.com/channels/%s/%s/%s\n", m.GuildID, m.ChannelID, m.ID)
	for i := 0; i < 3; i++ {
		msg += url
	}

	if secondsLeft > 0 {
		msg += fmt.Sprintf("Kegs kann noch %d Sekunden eingesammelt werden.", secondsLeft)
	} else {
		msg += "Kegs kann nicht mehr eingesammelt werden."
	}

	return msg
}

func sendPingMessage(s *discordgo.Session, m *discordgo.Message) (*discordgo.Message, error) {
	res, err := s.ChannelMessageSend(DISCORD_PING_CHANNEL_ID.Value(), pingMessage(m, int(countdown.Seconds())))
	if err != nil {
		return res, err
	}
	log.Printf("Responded to event system action within %s.", time.Since(m.Timestamp).Round(time.Millisecond))

	for i := int(countdown.Seconds()) - 1; i >= 0; i-- {
		now := time.Now()
		if _, err = s.ChannelMessageEdit(res.ChannelID, res.ID, pingMessage(m, i)); err != nil {
			log.Printf("Failed to edit message to update countdown: %v", err)
			continue
		}
		time.Sleep(time.Second - time.Since(now))
	}

	return res, nil
}
