package main

import (
	"io/ioutil"
	"os"

	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/BurntSushi/toml"
)

type Config struct {
	ListID int `toml:"list_id"`
}

func loadCurrentConfig() (c *Config, err error) {
	f, err := os.Open(".wlcli")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	_, err = toml.Decode(string(b), &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
