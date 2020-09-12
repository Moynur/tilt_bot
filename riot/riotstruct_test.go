package riot_test

import (
	"reflect"
	"testing"
	"tiltbot/riot"
)

var expectedStruct = riot.SummonerName{
	// ID is encrypted SUMMONER ID
	ID: "Lmm2gqd1xTqHZapAhz6upBqCqeIL6yyXuZ1sqlf2Lzlx7Hg",
	// AccountID is encrypted Account ID
	AccountID:     "ZAuc8LTu1rf3SnG08hewotljqtneFj1XgMFzhuJZczkVArM",
	Puuid:         "fXrHkmwdTpr-GJ1Pd7CIITlU4c-Gqmn79tUkux2F7EynBaK18jov31L4RpdUA_EgrJPj151Z8e7YFQ",
	Name:          "HepticHorror",
	ProfileIconID: 4496,
	RevisionDate:  1591971410,
	SummonerLevel: 218,
}

// This is the JSON format of the expected struct and vice versa
var inputJSON []byte = byte(`{"id":"Lmm2gqd1xTqHZapAhz6upBqCqeIL6yyXuZ1sqlf2Lzlx7Hg","accountId":"ZAuc8LTu1rf3SnG08hewotljqtneFj1XgMFzhuJZczkVArM","puuid":"fXrHkmwdTpr-GJ1Pd7CIITlU4c-Gqmn79tUkux2F7EynBaK18jov31L4RpdUA_EgrJPj151Z8e7YFQ","name":"HepticHorror","profileIconId":4496,"revisionDate":1591971410,"summonerLevel":218}`

// Need to refacor tests woo 
func TestConstructGood(t *testing.T) {
	actual, _ := riot.MaKeSummonerName(inputJSON)
	if !reflect.DeepEqual(actual, expectedStruct) {
		t.Fatalf("FAIL: \nExpected: %#v\nActual: %#v", expectedStruct, actual)
	}
	t.Logf("PASS")
}

func TestConstructFail(t *testing.T) {
	badJSON := "2131231"
	_, err := riotstruct.MaKeSummonerName(badJSON)
	if err == nil {
		t.Fatalf("Expected an error and didn't get one")
	}
	t.Logf("Pass")
}
