package commands

import (
	"DiscordBot/infra/db"
	"DiscordBot/infra/embedMessages"
	"DiscordBot/infra/entity"
	"fmt"
	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func DeletePhrase(ctx *disgolf.Ctx) {
	options := ctx.Options
	subject := options["creator"].UserValue(ctx.Session).Mention()
	phraseDel := options["phrase"].StringValue()

	database, err := db.OpenDb()
	defer db.CloseDB(database)
	if err != nil {
		return
	}
	var phrase entity.Phrase
	err = database.Delete(&phrase, "creator = ? AND phrase = ? AND guild_id = ?", subject, phraseDel, ctx.Interaction.GuildID).Error
	fmt.Println(phrase)
	if err != nil {
		//ctx.Reply(fmt.Sprintf("Não foi possivel deletar a frase **%s** de **%s**"), phraseDel, phraseDel)
		embed := embedMessages.NewEmbed().
			SetTitle("Error while deleting the Pearl Pérola").
			AddField(phraseDel, subject).
			SetColor(0xCC0000).MessageEmbed
		ctx.Session.ChannelMessageSendEmbed(ctx.Interaction.ChannelID, embed)
		return
	}

	embed := embedMessages.NewEmbed().
		SetTitle("Pérola successfully deleted").
		AddField(phraseDel, subject).
		SetColor(0x67339B).
		MessageEmbed

	ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{embed}},
	})

	//ctx.Session.ChannelMessageSendEmbed(ctx.Interaction.ChannelID, embed)
}
