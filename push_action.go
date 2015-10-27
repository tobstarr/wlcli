package main

import "strings"

type pushAction struct {
	Payload []string `cli:"arg required"`
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
	t := &task{ListID: ib.ID, Title: strings.Join(r.Payload, " ")}
	_, err = cl.CreateTask(t)
	return err
}
