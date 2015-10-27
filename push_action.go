package main

import "strings"

type pushAction struct {
	Payload []string `cli:"arg required"`
	Tags    []string `cli:"opt --tags"`
}

func (r *pushAction) Run() error {
	cl, err := loadClient()
	if err != nil {
		return err
	}
	ib, err := cl.Inbox()
	if err != nil {
		return err
	}
	parts := r.Payload
	for _, t := range r.Tags {
		parts = append(parts, "#"+t)
	}
	t := &task{ListID: ib.ID, Title: strings.Join(parts, " ")}
	_, err = cl.CreateTask(t)
	return err
}
