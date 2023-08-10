package routers

import (
	"DiscordBot/commands"
	"strings"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func addPerolaRouter(dc *disgolf.Bot) {
	addperolaCommand := &disgolf.Command{
		Name:        strings.ToLower("addPerola"),
		Description: "Adiciona uma nova Perola no servidor",
		Type:        discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			commands.AddPhrase(ctx)
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
	dc.Router.Register(addperolaCommand)
}
