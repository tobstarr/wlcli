package main

import (
	"strings"

	"github.com/tobstarr/wlcli/wlclient"
)

type pushAction struct {
	listID  int
	Payload []string `cli:"arg required"`
	Tags    []string `cli:"opt --tags"`
}

func (r *pushAction) Run() error {
	cl, err := loadClient()
	if err != nil {
		return err
	}
	listID := r.listID
	if listID == 0 {
		ib, err := cl.Inbox()
		if err != nil {
			return err
		}
		listID = ib.ID
	}
	parts := r.Payload
	for _, t := range r.Tags {
		parts = append(parts, "#"+t)
	}
	t := &wlclient.Task{ListID: listID, Title: strings.Join(parts, " ")}
	_, err = cl.CreateTask(t)
	return err
}
