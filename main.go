package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	//Loads DISCORD_TOKEN variable from .env
	godotenv.Load()
	token := os.Getenv("DISCORD_TOKEN")
	//Creates session
	dg, _ := discordgo.New("Bot " + token)

	//Sets up a listener
	dg.AddHandler(sendMessage)

	//Opens up a connection
	dg.Open()
	defer dg.Close()

	//Blocks from quitting unless any of the signals below are received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Checks to see if message author is bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Sends a message
	s.ChannelMessageSend(m.ChannelID, "Hello there")
}
