package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func newMemberHandler(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	fmt.Printf("Novo membro no servidor %s:\n", m.GuildID)
	fmt.Printf("ID do Membro: %s\n", m.Member.User.ID)
	fmt.Printf("Nome do Membro: %s\n", m.Member.User.Username)
}
