package handler

import (
	"DiscordBot/infra/handler/commands"
	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Routers(dc *disgolf.Bot) {

	pingCommand := &disgolf.Command{
		Name:        "ping",
		Description: "pong",
		Type:        discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			commands.PingPongHandler(ctx)
		}),
	}
	dc.Router.Register(pingCommand)

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
