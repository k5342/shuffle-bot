package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func createNTeams(nTeams int, users []string) ([][]string, error) {
	nActiveUsers := len(users)
	nMembers := int(math.Round(float64(nActiveUsers) / float64(nTeams)))

	if nMembers < 1 || nActiveUsers < nTeams {
		return nil, errors.New("more users required!!!")
	}

	// shuffle by connected users
	idx := rand.Perm(nActiveUsers)

	var shuffledUsers []string
	for _, newIdx := range idx {
		shuffledUsers = append(shuffledUsers, users[newIdx])
	}

	// devide into {nTeams} teams
	result := make([][]string, nTeams)
	for i := 0; i < nTeams-1; i++ {
		result[i] = shuffledUsers[i*nMembers : (i+1)*nMembers]
	}
	result[nTeams-1] = shuffledUsers[(nTeams-1)*nMembers : len(shuffledUsers)]
	fmt.Println(result)

	return result, nil
}
