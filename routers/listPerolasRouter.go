package routers

import (
	"DiscordBot/commands"
	"strings"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func listPerolasRouter(dc *disgolf.Bot) {
	listperolasCommand := &disgolf.Command{
		Name:        strings.ToLower("listPerolas"),
		Description: "Lista as Pérolas do servidor",
		Type:        discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			commands.ListPhrases(ctx)
		}),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionMentionable,
				Name:        "creator",
				Description: "Creator of the phrase",
				Required:    false,
			},
		},
	}
	dc.Router.Register(listperolasCommand)
}
