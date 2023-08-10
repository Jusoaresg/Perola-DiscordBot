package embedMessages

import "github.com/bwmarrin/discordgo"

func SuccessEmbed(s *discordgo.Session, channelID string, successMessage string) {
	success := NewEmbed().
		AddField("", successMessage).
		SetColor(0x00ff00).MessageEmbed

	s.ChannelMessageSendEmbed(channelID, success)
}
