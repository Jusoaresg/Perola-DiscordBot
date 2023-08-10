package routers

import (
	"DiscordBot/commands"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func pingRouter(dc *disgolf.Bot) {
	pingCommand := &disgolf.Command{
		Name:        "ping",
		Description: "pong",
		Type:        discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			commands.PingPongHandler(ctx)
		}),
	}
	dc.Router.Register(pingCommand)
}
