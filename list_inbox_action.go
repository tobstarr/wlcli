package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/dynport/gocli"
)

type listInboxAction struct {
	Tag string `cli:"opt --tag"`
}

func (r *listInboxAction) Run() error {
	cl, err := loadClient()
	if err != nil {
		return err
	}
	ib, err := cl.Inbox()
	if err != nil {
		return err
	}

	tasks, err := cl.Tasks(ListID(ib.ID))
	if err != nil {
		return err
	}
	sort.Sort(tasks)
	t := gocli.NewTable()
	for _, task := range tasks {
		if r.Tag != "" && !matchesTag(task.Title, r.Tag) {
			continue
		}
		t.Add(task.CreatedAt.Format("2006-02-01 15:04:05"), task.ID, task.Title)
	}
	fmt.Println(t)
	return nil
}

func matchesTag(text, tag string) bool {
	tag = strings.ToLower(tag)
	for _, f := range strings.Fields(strings.ToLower(text)) {
		if f == "#"+tag {
			return true
		}
	}
	return false
}
