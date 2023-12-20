package main

import "fmt"

func pingMessage(guildID, channelID, messageID string) string {
	msg := fmt.Sprintf("<@&%s>\n", DISCORD_PING_ROLE_ID.Value())

	url := fmt.Sprintf("https://discord.com/channels/%s/%s/%s\n", guildID, channelID, messageID)
	for i := 0; i < 3; i++ {
		msg += url
	}

	return msg
}
