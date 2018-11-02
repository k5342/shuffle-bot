package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// cache tables
var guildIDs map[string]string

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("SHUFFLEBOT_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord bot: ", err)
		return
	}

	guildIDs = make(map[string]string)
	dg.AddHandler(messageHandler)

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

func sendReply(s *discordgo.Session, m *discordgo.MessageCreate, str string) {
	sendMessage := fmt.Sprintf("<@!%s> ", m.Author.ID)
	sendMessage += str
	s.ChannelMessageSend(m.ChannelID, sendMessage)
}

func isContain(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
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

	guild, err := s.Guild(gid)
	if err != nil {
		fmt.Println("Error while fetching guild: ", err)
		return
	}

	// find users voice channel & fetch connected users
	voiceChannelUsers := map[string][]string{}
	var sourceVoiceChannel string
	for _, vs := range guild.VoiceStates {
		if vs.UserID == m.Author.ID {
			sourceVoiceChannel = vs.ChannelID
		}

		user, err := s.User(vs.UserID)
		if err != nil {
			fmt.Println("Error while fetching username")
			sendReply(s, m, "Error: unknown error.")
			return
		}
		if !isContain(user.Username, skipUsernames) {
			voiceChannelUsers[vs.ChannelID] =
				append(voiceChannelUsers[vs.ChannelID], user.Username)
		}
	}

	// not found in any voice channel
	if sourceVoiceChannel == "" {
		sendReply(s, m, "Please connect some voice channel!")
		return
	}

	// check nTeams
	totalUserCount := len(voiceChannelUsers[sourceVoiceChannel])

	nMembers := int(math.Round(float64(totalUserCount) / float64(nTeams)))
	if totalUserCount < nTeams {
		sendReply(s, m, fmt.Sprintf("More member required to make %d team(s) by %d member(s)!", nTeams, nMembers))
		return
	}

	// shuffle by connected users
	idx := rand.Perm(totalUserCount)

	var shuffledUsers []string
	for _, newIdx := range idx {
		shuffledUsers = append(shuffledUsers, voiceChannelUsers[sourceVoiceChannel][newIdx])
	}

	// devide into {nTeams} teams
	result := make([][]string, nTeams)
	for i := 0; i < nTeams-1; i++ {
		result[i] = shuffledUsers[i*nMembers : (i+1)*nMembers]
	}
	result[nTeams-1] = shuffledUsers[(nTeams-1)*nMembers : len(shuffledUsers)]
	fmt.Println(result)

	// send message
	outputString := fmt.Sprintf("created %d team(s)!\n", nTeams)
	for i := 0; i < nTeams; i++ {
		outputString += fmt.Sprintf("Team%d: %s\n", i+1, strings.Join(result[i], ", "))
	}
	sendReply(s, m, outputString)
}
