package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func userPresenceUpdateHandler(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	// update cache
	if p.User.Username != "" {
		fmt.Println("Username changed: " + p.User.Username)
		usernames[p.User.ID] = *p.User
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// translate channelID -> guildID to reduce latency
	// This does not need use in case of building with latest discordgo's develop branch
	gid, ok := guildIDs[m.ChannelID]
	if !ok {
		fmt.Println("Cache MISS")
		// cache miss
		sourceTextChannel, err := s.Channel(m.ChannelID)
		if err != nil {
			fmt.Println("Error while fetching source channel: ", err)
			return
		}
		gid = sourceTextChannel.GuildID
		guildIDs[m.ChannelID] = gid
	}

	if gid == "" {
		// Invoked from user chat directly
		s.ChannelMessageSend(m.ChannelID, "Please send after connecting and joining some voice channel!")
		return
	}

	// Invoked from Server (Guild)

	if !strings.HasPrefix(m.Content, "!!teams") {
		return
	}

	args := strings.Split(m.Content, " ")
	if len(args) <= 1 {
		sendReply(s, m, "Usage: `!!teams <number of teams to create> [skip username ...]`")
		return
	}

	var skipUsernames []string
	if len(args) > 2 {
		skipUsernames = args[2:len(args)]
	}

	_nTeams, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		fmt.Println("Error while parsing user specified value: ", err)
		sendReply(s, m, "Please specify in number!!!")
		return
	}

	nTeams := int(_nTeams)

	if nTeams <= 0 || nTeams >= 100 {
		if gid == "223518751650217994" {
			// for internal uses
			sendReply(s, m, "<:kakattekoi:461046115257679872>")
		} else {
			sendReply(s, m, "Please specify in *realistic* number!!!!!")
		}
		return
	}

	// find users voice channel & fetch connected users
	voiceChannelUsers, err := parseVoiceState(s, m.Author.ID, gid, skipUsernames)
	if err != nil {
		sendReply(s, m, err.Error())
		return
	}

	result, err := createNTeams(nTeams, voiceChannelUsers)
	if err != nil {
		sendReply(s, m, err.Error())
		return
	}

	// send message
	outputString := fmt.Sprintf("created %d team(s)!\n", nTeams)
	for i := 0; i < nTeams; i++ {
		outputString += fmt.Sprintf("Team%d: %s\n", i+1, strings.Join(result[i], ", "))
	}
	sendReply(s, m, outputString)
}
