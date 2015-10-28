package main

import (
	"encoding/json"
	"os"
)

type clientConfig struct {
	ClientID    string `json:"client_id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

func loadClientConfig() (c *clientConfig, err error) {
	f, err := os.Open(os.ExpandEnv("$HOME/.wlcli/config.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return c, json.NewDecoder(f).Decode(&c)
}
