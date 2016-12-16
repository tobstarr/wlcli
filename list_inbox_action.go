package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/dynport/gocli"
	"github.com/tobstarr/wlcli/wlclient"
)

type listInboxAction struct {
	listID int
	Tag    string `cli:"opt --tag"`
	Limit  int    `cli:"opt --limit default=25"`
	All    bool   `cli:"opt --all"`
	Full   bool   `cli:"opt --full"`
}

// this should be a generic inbox action
func (r *listInboxAction) Run() error {
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

	if dataOnStdin() {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		task, err := cl.CreateTask(&wlclient.Task{ListID: listID, Title: strings.TrimSpace(string(b))})
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(task)
		return nil
	}
	tasks, err := cl.Tasks(wlclient.ListID(listID))
	if err != nil {
		return err
	}

	pos, err := cl.TaskPositions(listID)
	if err != nil {
		return err
	}
	if len(pos) == 0 {
		return fmt.Errorf("unable to get positions for list %d", listID)
	}

	if len(tasks) == 0 {
		fmt.Printf("no tasks found. use `wlcli push` to create one\n")
		return nil
	}
	sort.Sort(tasks)
	t := gocli.NewTable()
	tm := map[int]*wlclient.Task{}
	for _, t := range tasks {
		tm[t.ID] = t
	}

	positions := pos[0].Values
	for i, p := range positions {
		task, ok := tm[p]
		if !ok {
			continue
		}
		delete(tm, p)
		if r.Tag != "" && !matchesTag(task.Title, r.Tag) {
			continue
		}
		title := task.Title
		if !r.Full {
			title = truncate(title, 64)
		}
		if i < r.Limit && !r.All {
			t.Add(task.ID, title)
		}
	}
	for _, task := range tm {
		if r.Tag != "" && !matchesTag(task.Title, r.Tag) {
			continue
		}
		t.Add(task.ID, truncate(task.Title, 64))
	}
	fmt.Println(t)
	return nil
}

func dataOnStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// this is naive!!!
func truncate(in string, max int) string {
	parts := strings.Split(in, "")
	if len(parts) < max {
		return in
	}
	return strings.Join(parts[0:max], "")
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
