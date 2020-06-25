package main

import (
	"fmt"
	"os"
)

var (
	riotKey string
)

type config struct {
	riotKey string
}

// Check your vars exist
func check() (*config, error) {
	for k, v := range map[string]*string{
		"riot_key": &riotKey,
	} {
		var ok bool
		if *v, ok = os.LookupEnv(k); !ok {
			return nil, fmt.Errorf("missing environment variable %s", k)
		}
	}
	return &config{
		riotKey: riotKey,
	}, nil
}
