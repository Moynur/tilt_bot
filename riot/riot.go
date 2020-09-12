package riot

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetMatchHistory contains the fields needed to query RIOTs API to see someones match info
type GetMatchHistory struct {
	encryptedID string
	key         string
	name        string
}

var nameURL string = "https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name"
var matchURL string = "https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner"

// New stores info so we don't get messy repeatedly entering these, should revisit error handling at some point, note we don't store encryptedID yet as we'll need to query this
func New(name, key string) (*GetMatchHistory, *SummonerName, error) {

	if len(key) == 0 || len(key) == 0 {
		return nil, nil, fmt.Errorf("One of your variables is empty, \nname:%v \nkey:%v ", name, key)
	}
	respBody, err := GetInfo(nameURL, name, key)
	if err != nil {
		return nil, nil, fmt.Errorf("Somethings gone wrong: %v", err)
	}
	newStruct, err := MaKeSummonerName(respBody)

	if err != nil {
		return nil, nil, fmt.Errorf("Couldn't build struct after polling api %v", err)
	}
	return &GetMatchHistory{name: name, key: key, encryptedID: newStruct.ID}, &newStruct, nil
}

// GetInfo will be used to construct all http requests
func GetInfo(BaseURL, QueryParam, authKey string) ([]byte, error) {
	if authKey == "" {
		return nil, errors.New("No Auth Header")
	}
	NewURL := fmt.Sprintf("%v/%v", BaseURL, QueryParam)
	maxTime := time.Second * 10
	client := &http.Client{Timeout: maxTime}
	req, err := http.NewRequest("GET", NewURL, nil)
	if err != nil {
		return nil, errors.New("Unable to set request 400")
	}
	req.Header.Set("X-Riot-Token", authKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("server didn’t respond 200 OK: %v", resp.StatusCode)
	}
	BodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, errors.New("Server error: 500")
	}
	return BodyText, nil

}

// GetSummonerByName function will lookup summoner by name
// Can refactor later to remove URL from function not entirely necessary

// GetSummonerMatchHistory uses that information to perform a single check on summoner
func (s *GetMatchHistory) GetSummonerMatchHistory() (FlexQ, SoloQ AccountGamesPlayed, err error) {
	resp, err := GetInfo(matchURL, s.encryptedID, s.key)
	if err != nil {
		return FlexQ, SoloQ, fmt.Errorf("Somethings gone wrong: %v", err)
	}
	FlexQ, SoloQ, err = MakeGamesPlayed(resp)
	return FlexQ, SoloQ, nil
}

// Poll essentially the same as Request but this will poll rather than perform a single request
func (s *GetMatchHistory) Poll(c chan AccountGamesPlayed, tickrate int) {
	ticker := time.NewTicker(time.Second * time.Duration(tickrate)).C
	for {
		select {
		case <-ticker:
			GetInfo(matchURL, s.encryptedID, s.key)
			fmt.Println("I just did a request")
			resp, err := GetInfo(matchURL, s.encryptedID, s.key)
			if err != nil {
				panic("implement error handler plz")
			}
			_, SoloQ, err := MakeGamesPlayed(resp)
			fmt.Println("increased losses by 1 ", SoloQ.Losses)
			c <- SoloQ
		}

	}
}
