package commands

import (
	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func PingPongHandler(ctx *disgolf.Ctx) {

	ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		}})
	return

}
