package commands

import (
	"DiscordBot/infra/db"
	"DiscordBot/infra/embedMessages"
	"DiscordBot/infra/entity"
	"fmt"
	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

var message *discordgo.Message // Variável global para armazenar a mensagem embed
var currentPage int            // Variável global para acompanhar a página atual
var pageCount int              // Variável global para acompanhar o número total de páginas
var phrases []entity.Phrase    // Variável global para armazenar as frases
var perPage int                // Variável global para armazenar a quantidade de frases por página

func ListPhrases(ctx *disgolf.Ctx) {
	database, err := db.OpenDb()
	if err != nil {
		return
	}

	if ctx.Options["creator"] != nil {

		database.Find(&phrases, "guild_id = ? AND creator = ?", ctx.Interaction.GuildID, ctx.Options["creator"].UserValue(ctx.Session).Mention())
	} else {
		database.Find(&phrases, "guild_id = ?", ctx.Interaction.GuildID)
	}

	perPage = 10 // Quantidade de frases por página
	pageCount = (len(phrases) + perPage - 1) / perPage

	ctx.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "Aqui estão as suas Pérolas."},
	})

	// Envia a primeira página
	currentPage = 0
	embed := createEmbed(currentPage)
	message, err = ctx.Session.ChannelMessageSendEmbed(ctx.Interaction.ChannelID, embed) // Atribui à variável global
	if err != nil {
		return
	}

	// Adiciona as reações para navegar entre as páginas
	err = ctx.Session.MessageReactionAdd(ctx.Interaction.ChannelID, message.ID, "⬅️")
	if err != nil {
		return
	}
	err = ctx.Session.MessageReactionAdd(ctx.Interaction.ChannelID, message.ID, "➡️")
	if err != nil {
		return
	}

	// Função para tratar eventos de reação
	ctx.Session.AddHandler(handleReaction)
}

func createEmbed(page int) *discordgo.MessageEmbed {
	start := page * perPage
	end := start + perPage
	if end > len(phrases) {
		end = len(phrases)
	}

	embed := embedMessages.NewEmbed().
		SetTitle(fmt.Sprintf("Pérolas (Página %d/%d)", page+1, pageCount)) // Exibe o número da página atual

	for i := start; i < end; i++ {
		msg := fmt.Sprintf("**```%s``` de %s** [salva em %s por %s]",
			phrases[i].Phrase,
			phrases[i].Creator,
			phrases[i].CreatedAt.Format("02-01-2006 as 3:04PM"),
			phrases[i].SavedBy)
		embed.AddField("", msg)
	}

	return embed.MessageEmbed
}

func handleReaction(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return // Ignora as reações do próprio bot
	}

	if message == nil || r.MessageID != message.ID { // Verifica se message é nulo ou se os IDs não correspondem
		return
	}

	switch r.Emoji.Name {
	case "⬅️":
		s.MessageReactionRemove(r.ChannelID, r.MessageID, "⬅️", r.UserID)
		if currentPage > 0 {
			currentPage--
		}
	case "➡️":
		s.MessageReactionRemove(r.ChannelID, r.MessageID, "➡️", r.UserID)
		if currentPage < pageCount-1 {
			currentPage++
		}
	}

	// Atualiza a mensagem embed com base na página atual
	embed := createEmbed(currentPage)
	_, err := s.ChannelMessageEditEmbed(r.ChannelID, message.ID, embed)
	if err != nil {
		return
	}
}
