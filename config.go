package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/BurntSushi/toml"
)

const configFileName = ".wlcli"

type Config struct {
	ListID int `toml:"list_id"`
}

func loadCurrentConfig() (c *Config, err error) {
	current, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	for current != "" {
		path := filepath.Join(current, configFileName)
		dbg.Printf("testing path %s", path)
		c, err := loadConfig(path)
		if err != nil {
			newCurrent := filepath.Dir(current)
			if newCurrent == current {
				break
			}
			current = newCurrent
		} else {
			dbg.Printf("found path %s", path)
			return c, nil
		}
	}
	return nil, nil
}

func loadConfig(path string) (c *Config, err error) {
	f, err := os.Open(path)
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
