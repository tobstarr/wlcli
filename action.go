package main

import (
	"bufio"
	"io"
	"strings"
)

type rawAction struct {
	Action      string
	ID          string
	Message     string
	Description string
}

func extractRawActions(in string) (out []*rawAction, err error) {
	r := bufio.NewReader(strings.NewReader(in))
	var current *rawAction
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		line = strings.TrimSpace(line)
		if ra := extractRawAction(line); ra != nil {
			if current != nil {
				current.Description = strings.TrimSpace(current.Description)
			}
			out = append(out, ra)
			current = ra
		} else if current != nil {
			if isComment(line) {
				continue
			}
			if len(current.Description) > 0 {
				current.Description += "\n"
			}
			current.Description += line
		}
	}
	if current != nil {
		current.Description = strings.TrimSpace(current.Description)
	}
	return out, nil
}

func extractRawAction(in string) (out *rawAction) {
	fields := strings.Fields(in)
	if len(fields) < 3 {
		return nil
	}

	if len(fields[0]) != 1 { // replace this with something smarter
		return nil
	}
	return &rawAction{Action: fields[0], ID: fields[1], Message: strings.Join(fields[2:], " ")}
}
