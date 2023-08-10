package routers

import (
	"DiscordBot/commands"
	"strings"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func deletePerolaRouter(dc *disgolf.Bot) {
	deleteperolaCommand := &disgolf.Command{
		Name:        strings.ToLower("deletePerola"),
		Description: "Deleta uma perola",
		Type:        discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			commands.DeletePhrase(ctx)
		}),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionMentionable,
				Name:        "creator",
				Description: "Creator of the phrase",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "phrase",
				Description: "Phrase to create",
				Required:    true,
			},
		},
	}
	dc.Router.Register(deleteperolaCommand)
}
