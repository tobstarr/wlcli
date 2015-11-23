package main

import (
	"log"
	"os"

	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/dynport/dgtk/cli"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	router := cli.NewRouter()

	c, _ := loadCurrentConfig()
	if c != nil && c.ListID > 0 {
		router.Register("push", &pushAction{listID: c.ListID}, "Push a task to the current list")
		router.Register("list", &listInboxAction{listID: c.ListID}, "List current list")
		router.Register("edit", &edit{listID: c.ListID}, "Edit list")
	} else {
		router.Register("edit", &edit{}, "Edit list")
	}
	router.Register("inbox/list", &listInboxAction{}, "List Inbox")
	router.Register("inbox/push", &pushAction{}, "Push a task to inbox")
	router.Register("init", &initAction{}, "Initialize "+configFileName+" config file")
	router.Register("lists/list", &listLists{}, "List Lists")
	router.Register("tasks/complete", &completeTasks{}, "Complete Tasks")
	router.Register("tasks/delete", &deleteTasks{}, "Delete Tasks")
	switch err := router.RunWithArgs(); err {
	case nil, cli.ErrorHelpRequested, cli.ErrorNoRoute:
		// ignore
		return
	default:
		l.Fatal(err)
	}
}
