package main

import (
	"fmt"
	"log"
	"tiltbot/riot"
)

var (
	name = "heptichorror"
)

func main() {
	// Look to build annoy package maybe
	// In the end we want to be able to annoy multiple people, ie go Annoy("summoner") instead of just Annoy
	vars, err := check()
	if err != nil {
		log.Fatalf("Couldn't load vars, got error: %v", err)
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
	go riotService.Poll(10)
	for {
	}
}
