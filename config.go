package main

import (
	"fmt"
	"os"
)

var (
	riotKey          string
	discordToken     string
	discordChannelID string
)

type config struct {
	riotKey          string
	discordToken     string
	discordChannelID string
}

// Check your vars exist
func check() (*config, error) {
	for k, v := range map[string]*string{
		"riot_key":             &riotKey,
		"discord_server_token": &discordToken,
		"TargetChannelID2":     &discordChannelID,
	} {
		var ok bool
		if *v, ok = os.LookupEnv(k); !ok {
			return nil, fmt.Errorf("missing environment variable %s", k)
		}
	}
	return &config{
		riotKey:          riotKey,
		discordToken:     discordToken,
		discordChannelID: discordChannelID,
	}, nil
}
