package main

import "testing"

func TestExtractRawActions(t *testing.T) {
	actions, err := extractRawActions(re)
	if err != nil {
		t.Fatal(err)
	}
	if v, ex := len(actions), 3; ex != v {
		t.Fatalf("expected to get %d actions, was %d", ex, v)
	}
	if v, ex := actions[0].Action, "p"; ex != v {
		t.Errorf("expected first action to be %q, was %q", ex, v)
	}
	if v, ex := actions[0].ID, "1234"; ex != v {
		t.Errorf("expected first id to be %q, was %q", ex, v)
	}
	if v, ex := actions[0].Message, "first message"; ex != v {
		t.Errorf("expected message to be %q, was %q", ex, v)
	}
	if v, ex := actions[1].Description, "description\n\ntext 1"; ex != v {
		t.Errorf("expected description of to be %q, was %q", ex, v)
	}
}

const re = `p 1234 first message
p 1235 second task

description

text 1

p 1236 third task
`
