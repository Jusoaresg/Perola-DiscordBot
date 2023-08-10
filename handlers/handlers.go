package handlers

import (
	"DiscordBot/routers"
	"DiscordBot/services/db"
	"fmt"
	"log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func Handlers(dc *disgolf.Bot) {

	dc.AddHandler(dc.Router.HandleInteraction)

	dc.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is running!")
	})

	// Bot enter server
	dc.AddHandler(func(s *discordgo.Session, g *discordgo.GuildCreate) {
		database, err := db.InitDB()
		defer db.CloseDB(database)
		guildID := g.ID
		serverName := g.Name

		err = dc.Router.Sync(dc.Session, "", g.ID)

		_, err = db.SearchIfGuildExist(database, guildID, serverName)
		if err != nil {
			fmt.Println("Error on guildCreate ", err)
		}
		fmt.Printf("Bot entrou no servidor {%s} com ID de {%s}\n", g.Name, g.ID)
	})

	routers.Routers(dc)
}
