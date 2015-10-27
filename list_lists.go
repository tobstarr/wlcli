package main

import (
	"fmt"

	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/dynport/gocli"
)

type listLists struct {
}

func (r *listLists) Run() error {
	c, err := loadClient()
	if err != nil {
		return err
	}
	lists, err := c.Lists()
	if err != nil {
		return err
	}
	t := gocli.NewTable()
	for _, list := range lists {
		t.Add(list.ID, list.Title)
	}
	fmt.Println(t)
	return nil
}
