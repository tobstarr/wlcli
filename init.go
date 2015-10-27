package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/BurntSushi/toml"
	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/dynport/gocli"
)

type initAction struct {
}

func (r *initAction) Run() error {
	path := configFileName
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("config file %s already exists", path)
	}
	cl, err := loadClient()
	if err != nil {
		return err
	}
	lists, err := cl.Lists()
	if err != nil {
		return err
	}
	t := gocli.NewTable()

	for _, l := range lists {
		t.Add(l.Title)
	}
	t.SortBy = 0
	sort.Sort(t)
	list := lists[t.Select("select a list")]
	l.Printf("initializing list %d %s", list.ID, list.Title)
	c := &Config{ListID: list.ID}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(c)
}
