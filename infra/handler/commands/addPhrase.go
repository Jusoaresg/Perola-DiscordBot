package commands

import (
	"DiscordBot/infra/embedMessages"
	"DiscordBot/infra/entity"
	"fmt"
	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func AddPhrase(ctx *disgolf.Ctx) {

	options := ctx.Options
	autor := options["creator"].UserValue(ctx.Session).Mention()
	phrase := options["phrase"].StringValue()

	db, err := gorm.Open(sqlite.Open("sql.db"), &gorm.Config{})
	if err != nil {
		return
	}

	var guild entity.GuildEntity
	err = db.Where("guild_id = ?", ctx.Interaction.GuildID).Preload("Phrases").First(&guild).Error
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

	guild.Phrases = append(guild.Phrases, frase)
	err = db.Save(&guild.Phrases).Error
	if err != nil {
		fmt.Println("Erro ao adicionar a nova frase:", err)
		embedMessages.ErrorEmbedMessage(ctx.Session, ctx.State.SessionID, "Erro ao adicionar a nova frase")
		return
	}

	//Mensagem de criação de mensagem
	fmt.Println("Nova frase adicionada com sucesso!")
	embed := embedMessages.NewEmbed().
		SetTitle("Pérola adicionada").
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
