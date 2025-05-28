package main

import (
	"discord-finance/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env not found")
	}
	Token := os.Getenv("TOKEN")
	if Token == "" {
		fmt.Println("token not found")
		return
	}
	fmt.Println("token is: ", Token)

	config.InitSupabase()

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("err create session", err)
		return
	}
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("err open connection,", err)
		return
	}
	fmt.Println("bot is online")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("[DEBUG] Message received: %s from %s\n", m.Content, m.Author.Username)

	switch m.Content {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "pong")
	case "saldo":
		s.ChannelMessageSend(m.ChannelID, "saldo")
	}
}
