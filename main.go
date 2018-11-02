package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// cache tables
var (
	guildIDs  map[string]string
	usernames map[string]discordgo.User
)

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("SHUFFLEBOT_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord bot: ", err)
		return
	}

	// create cache
	guildIDs = make(map[string]string)
	usernames = make(map[string]discordgo.User)

	dg.AddHandler(messageHandler)
	dg.AddHandler(userPresenceUpdateHandler)

	dg.Open()
	if err != nil {
		fmt.Println("Error opening WebSocket connection: ", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
