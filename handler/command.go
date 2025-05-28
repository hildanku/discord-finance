package handler

import (
	"discord-finance/entity"
	"discord-finance/repository"
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handleIncomeCommand(s *discordgo.Session, m *discordgo.MessageCreate, parts []string) {
	if len(parts) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Format: `!income <amount>`")
		return
	}
	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "amount must be a number")
		return
	}

	tx := entity.Transaction{
		UserID:    m.Author.ID,
		Type:      "income",
		Amount:    amount,
		Note:      "xxxx",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	if err := repository.StoreTransaction(tx); err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed to record income.")
		fmt.Println("[ERROR] StoreTransaction:", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Income of %d has been recorded", amount))
}

func handleExpenseCommand(s *discordgo.Session, m *discordgo.MessageCreate, parts []string) {
	if len(parts) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Format: `!expense <amount>`")
		return
	}
	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "amount must be a number")
		return
	}

	tx := entity.Transaction{
		UserID:    m.Author.ID,
		Type:      "expense",
		Amount:    amount,
		Note:      "",
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := repository.StoreTransaction(tx); err != nil {
		s.ChannelMessageSend(m.ChannelID, "failed to record expense")
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("expense of %d has vbeen recorded!", amount))
}

func handleBalanceCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	balance, err := repository.GetBalanceByUserID(m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "failed to get balance")
		return
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("your balance is: Rp %d", balance))
}
