package main

import (
	"fmt"
	"log"
	"tiltbot/messagehandler"
	"tiltbot/riot"
)

var (
	name     = "heptichorror"
	tickrate = 10
)

func main() {
	// Look to build annoy package maybe
	// In the end we want to be able to annoy multiple people, ie go Annoy("summoner") instead of just Annoy
	vars, err := check()
	if err != nil {
		log.Fatalf("Couldn't load vars, got error: %v", err)
	}
	discordSessh, err := messagehandler.NewConnection(vars.discordToken)
	if err != nil {
		log.Fatalf("Something is probably empty, %v", err)
	}
	riotService, _, err := riot.New(name, vars.riotKey)
	if err != nil {
		log.Fatalf("Something is probably empty, %v", err)
	}
	FlexQ, SoloQ, err := riotService.GetSummonerMatchHistory()
	if err != nil {
		log.Fatalf("Request failed, %v", err)
	}
	fmt.Println("This should be Soloq: ", SoloQ.QueueType, SoloQ.LeaguePoints)
	fmt.Println("This should be FlexQ: ", FlexQ.QueueType, SoloQ.LeaguePoints)
	testStruct := make(chan riot.AccountGamesPlayed)
	go riotService.Poll(testStruct, tickrate)

	for {
		fmt.Println("start loop")
		select {
		case UpdatedSoloq := <-testStruct:
			if SoloQ.Losses > UpdatedSoloq.Losses {
				fmt.Println("triggered")
				discordSessh.SendMessage(vars.discordChannelID, "You guys suck")
				discordSessh.Close()
			}
			if SoloQ.Wins < UpdatedSoloq.Wins {
				discordSessh.SendMessage(vars.discordChannelID, "You guys may have won but you still suck")
				discordSessh.Close()
			}

		}
	}
}
