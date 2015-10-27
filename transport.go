package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type transport struct {
	ClientID    string `json:"client_id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	if t.ClientID == "" || t.AccessToken == "" {
		return nil, fmt.Errorf("ClientID and AccessToken must both be set")
	}
	r.Header.Set("X-Client-ID", t.ClientID)
	r.Header.Set("X-Access-Token", t.AccessToken)
	return http.DefaultClient.Do(r)
}

func loadTransport() (c *transport, err error) {
	f, err := os.Open(os.ExpandEnv("$HOME/.wundercli/config.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return c, json.NewDecoder(f).Decode(&c)
}
