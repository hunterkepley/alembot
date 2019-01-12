package main

import (
	"github.com/bwmarrin/discordgo"
)

func alembicCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, ":alembic:")
}
