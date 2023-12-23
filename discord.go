package koerbismaster

import "github.com/bwmarrin/discordgo"

func NewDiscordClient() (*discordgo.Session, error) {
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
