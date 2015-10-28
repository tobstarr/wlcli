package main

import "github.com/tobstarr/wlcli/wlclient"

func loadClient() (*wlclient.Client, error) {
	c, err := loadClientConfig()
	if err != nil {
		return nil, err
	}
	return wlclient.New(c.ClientID, c.AccessToken)
}
