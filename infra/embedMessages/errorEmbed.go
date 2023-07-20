package embedMessages

import (
	"github.com/bwmarrin/discordgo"
)

func ErrorEmbedMessage(s *discordgo.Session, channelID string, errorMessage string) {
	error := NewEmbed().
		SetColor(0xCC0000).
		AddField("", errorMessage).MessageEmbed
	s.ChannelMessageSendEmbed(channelID, error)
}
