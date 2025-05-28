package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	parts := strings.Fields(m.Content)
	if len(parts) == 0 {
		return
	}

	cmd := parts[0]
	switch cmd {
	case "!income":
		handleIncomeCommand(s, m, parts)
	case "!expense":
		handleExpenseCommand(s, m, parts)
	case "!saldo":
		handleBalanceCommand(s, m)
	}
}
