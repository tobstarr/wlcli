package main

import (
	"log"
	"os"

	"github.com/tobstarr/wlcli/Godeps/_workspace/src/github.com/dynport/dgtk/cli"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	router := cli.NewRouter()
	router.Register("push", &pushAction{}, "Push a task to inbox")
	switch err := router.RunWithArgs(); err {
	case nil, cli.ErrorHelpRequested, cli.ErrorNoRoute:
		// ignore
		return
	default:
		l.Fatal(err)
	}
}
