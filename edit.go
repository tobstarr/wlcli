package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tobstarr/wlcli/wlclient"
)

type edit struct {
	listID int
}

func (r *edit) Run() error {
	cl, err := loadClient()
	if err != nil {
		return err
	}
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return fmt.Errorf("EDITOR must be set in ENV")
	}
	tasks, err := cl.Tasks(wlclient.ListID(r.listID))
	if err != nil {
		return err
	}
	list, err := cl.List(r.listID)
	if err != nil {
		return err
	}
	tmpPath, dir, err := writeTasksToTempFile(list, tasks)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	c := exec.Command(editor, tmpPath)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	if err := c.Run(); err != nil {
		return err
	}
	f, err := os.Open(tmpPath)
	if err != nil {
		return err
	}
	actions, err := extractActions(cl, r.listID, f)
	f.Close()
	if err != nil {
		return err
	}
	for _, a := range actions {
		if err := a(); err != nil {
			return err
		}
	}
	return nil
}

type action func() error

type actions []action

var l = log.New(os.Stderr, "", 0)

func extractActions(cl *wlclient.Client, listID int, in io.Reader) (out actions, err error) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {

		action, err := extractActionFromLine(cl, listID, scanner.Text())
		if err != nil {
			l.Printf("WARN: %s", err)
			continue
		} else if action != nil {
			out = append(out, action)
		}
	}
	return out, scanner.Err()
}

func extractActionFromLine(cl *wlclient.Client, listID int, line string) (action, error) {
	if strings.HasPrefix(line, "# ") {
		return nil, nil
	}
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return nil, nil
	}
	if fields[0] == "*" && len(fields) > 1 {
		return func() error {
			_, err := cl.CreateTask(&wlclient.Task{ListID: listID, Title: strings.Join(fields[1:], " ")})
			return err
		}, nil
	}
	if len(fields) < 3 {
		return nil, nil
	}
	action, idString := fields[0], fields[1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, fmt.Errorf("%q does not parse to int: %s", idString, err)
	}

	switch action {
	case "r", "reword":
		return func() error {
			txt := strings.Join(fields[2:], " ")
			dbg.Printf("updating %d to %q", id, txt)
			return cl.UpdateTask(&wlclient.UpdateTask{ID: id, Title: s2p(txt)})
		}, nil
	case "c", "complete":
		return func() error { return cl.CompleteTask(id) }, nil
	case "d", "delete":
		return func() error { return cl.DeleteTask(id) }, nil
	case "k", "keep":
		return nil, nil
	default:
		return nil, fmt.Errorf("action %q not supported", action)
	}
}

func writeTasksToTempFile(list *wlclient.List, tasks wlclient.Tasks) (path, dir string, err error) {
	buf := &bytes.Buffer{}
	for _, t := range tasks {
		if _, err := fmt.Fprintf(buf, "k %d %s\n", t.ID, t.Title); err != nil {
			return "", "", err
		}
	}
	if len(tasks) > 0 {
		if _, err := io.WriteString(buf, "\n"); err != nil {
			return "", "", err
		}
	}
	if _, err := fmt.Fprintf(buf, "# Edit list %d: %s\n", list.ID, list.Title); err != nil {
		return "", "", err
	}
	if _, err := io.WriteString(buf, usage); err != nil {
		return "", "", err
	}
	d, err := ioutil.TempDir("/tmp", "wlcli-")
	if err != nil {
		return "", "", err
	}
	f, err := os.Create(filepath.Join(d, "git-rebase-todo"))
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	_, err = io.Copy(f, buf)
	if err != nil {
		return "", "", err
	}
	return f.Name(), d, nil
}

const usage = `# Commands
# k, keep = leave unchanged
# r, reword = use task, but edit the title
# c, complete = complete task
# d, delete = delete task
#
# to create create new tasks, add them on the bottom of the list:
# * task 1
# * task 2
`
