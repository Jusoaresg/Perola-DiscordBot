package commands

import (
	entity "DiscordBot/models"
	"DiscordBot/services/db"
	"DiscordBot/services/embedMessages"
	"fmt"
	"time"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func AddPhrase(ctx *disgolf.Ctx) {

	options := ctx.Options
	autor := options["creator"].UserValue(ctx.Session).Mention()
	phrase := options["phrase"].StringValue()

	database, err := db.OpenDb()
	if err != nil {
		return
	}
	defer db.CloseDB(database)

	var guild entity.GuildEntity
	err = database.Where("guild_id=?", ctx.Interaction.GuildID).Preload("Phrases").First(&guild).Error
	if err != nil {
		fmt.Println("Erro ao recuperar o GuildEntity:", err)
		return
	}

	frase := entity.Phrase{
		GuildID:   ctx.Interaction.GuildID,
		Creator:   autor,
		Phrase:    phrase,
		SavedBy:   ctx.Interaction.Member.Mention(),
		CreatedAt: time.Now(),
	}

	err = database.Create(&frase).Error
	if err != nil {
		fmt.Println("Erro ao criar frase:", err)
		return
	}

	guild.Phrases = append(guild.Phrases, frase)
	err = database.Save(&guild).Error
	if err != nil {
		fmt.Println("Erro ao adicionar a nova frase:", err)
		embedMessages.ErrorEmbedMessage(ctx.Session, ctx.State.SessionID, "Error adding new phrase")
		return
	}

	//Mensagem de criação de mensagem
	fmt.Println("Nova frase adicionada com sucesso!")
	embed := embedMessages.NewEmbed().
		SetTitle("Pérola added").
		SetColor(0x00ff00).
		AddField(phrase, autor).MessageEmbed
	ctx.Session.ChannelMessageSendEmbed(ctx.State.SessionID, embed)

	ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
