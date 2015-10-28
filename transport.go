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
	client      *client
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	if t.ClientID == "" || t.AccessToken == "" {
		return nil, fmt.Errorf("ClientID and AccessToken must both be set")
	}
	r.Header.Set("X-Client-ID", t.ClientID)
	r.Header.Set("X-Access-Token", t.AccessToken)
	return t.httpClient().Do(r)
}

func (t *transport) httpClient() *http.Client {
	if t.client != nil && t.client.Client != nil {
		return t.client.Client
	}
	return http.DefaultClient
}

func loadTransport() (c *transport, err error) {
	f, err := os.Open(os.ExpandEnv("$HOME/.wlcli/config.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return c, json.NewDecoder(f).Decode(&c)
}
