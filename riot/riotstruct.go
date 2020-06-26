package riot

import (
	"encoding/json"
)

// SummonerName struct
type SummonerName struct {
	// Summoner ID
	ID            string `json:"id"`
	AccountID     string `json:"accountID"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconID"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

// MaKeSummonerName builds a struct from the input into the SummonerName struct
func MaKeSummonerName(input []byte) (SummonerName, error) {
	var reading SummonerName
	err := json.Unmarshal([]byte(input), &reading)
	if err != nil {
		return reading, err
	}
	return reading, nil
}

// AccountGamesPlayed will return information on all ranked queues
type AccountGamesPlayed struct {
	SummonerNamestring string `json:"summonerName"`
	QueueType          string `json:"queueType"`
	LeaguePoints       int    `json:"leaguePoints"`
	Wins               int    `json:"wins"`
	Losses             int    `json:"losses"`
	HotStreak          bool   `json:"hotStreak"`
}

// MakeGamesPlayed builds a struct from the json input into the AccountGamesPlayed struct
func MakeGamesPlayed(input []byte) (FlexQ AccountGamesPlayed, SoloQ AccountGamesPlayed, err error) {
	var reading []AccountGamesPlayed
	err = json.Unmarshal([]byte(input), &reading)
	if reading[0].QueueType == "RANKED_FLEX_SR" { // look back at this later, consider looping over slice
		FlexQ = reading[0]
	} else if reading[0].QueueType == "RANKED_SOLO_5x5" {
		SoloQ = reading[0]

	}
	if reading[1].QueueType == "RANKED_SOLO_5x5" {
		SoloQ = reading[1]
	} else if reading[1].QueueType == "RANKED_FLEX_SR" {
		FlexQ = reading[1]

	}
	if err != nil {
		return FlexQ, SoloQ, err
	}
	return FlexQ, SoloQ, nil
}
