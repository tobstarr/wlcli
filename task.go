package main

import (
	"bytes"
	"encoding/json"
	"time"
)

type task struct {
	ID              int        `json:"id,omitempty"`
	ListID          int        `json:"list_id,omitempty"`
	Title           string     `json:"title,omitempty"`
	CreatedByID     int        `json:"created_by_id,omitempty"`
	Revision        int        `json:"revision,omitempty"`
	AssigneID       *int       `json:"assigne_id,omitempty"`
	Completed       *bool      `json:"completed,omitempty"`
	RecurrenceType  *string    `json:"recurrence_type,omitempty"`
	RecurrenceCount *int       `json:"recurrence_count,omitempty"`
	DueDate         *time.Time `json:"due_date,omitempty"`
	Starred         *bool      `json:"starred,omitempty"`
}

func (c *client) CreateTask(in *task) (out *task, err error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	return out, c.load("POST", "tasks", bytes.NewReader(b), &out)
}
