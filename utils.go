package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func isContain(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func sendReply(s *discordgo.Session, m *discordgo.MessageCreate, str string) {
	sendMessage := fmt.Sprintf("<@!%s> ", m.Author.ID)
	sendMessage += str
	s.ChannelMessageSend(m.ChannelID, sendMessage)
}

func parseVoiceState(s *discordgo.Session, uid string, gid string, skipUsernames []string) ([]string, error) {
	guild, err := s.Guild(gid)
	if err != nil {
		return nil, fmt.Errorf("error while fetching guild: %s", err)
	}

	voiceChannelUsers := map[string][]string{}
	var sourceVoiceChannel string
	for _, vs := range guild.VoiceStates {
		if vs.UserID == uid {
			sourceVoiceChannel = vs.ChannelID
		}

		// check cache
		user, ok := usernames[vs.UserID]
		if !ok {
			// cache MISS
			u, err := s.User(vs.UserID)
			if err != nil {
				return nil, fmt.Errorf("error while fetching username: %s", err)
			}
			user = *u
			usernames[vs.UserID] = user
		}

		if !isContain(user.Username, skipUsernames) {
			voiceChannelUsers[vs.ChannelID] =
				append(voiceChannelUsers[vs.ChannelID], user.Username)
		}
	}

	if sourceVoiceChannel == "" {
		return nil, fmt.Errorf("Please connect some voice channel!!!")
	} else {
		return voiceChannelUsers[sourceVoiceChannel], nil
	}
}
