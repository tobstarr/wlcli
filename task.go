package main

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

type task struct {
	ID              int        `json:"id,omitempty"`
	ListID          int        `json:"list_id,omitempty"`
	Title           string     `json:"title,omitempty"`
	CreatedByID     int        `json:"created_by_id,omitempty"`
	Revision        int        `json:"revision,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
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

type TaskOpt struct {
	ListID *int
}

func ListID(id int) func(*TaskOpt) {
	return func(o *TaskOpt) {
		o.ListID = &id
	}
}

func (t *TaskOpt) Encode() string {
	v := url.Values{}
	if t.ListID != nil {
		v.Add("list_id", strconv.Itoa(*t.ListID))
	}
	return v.Encode()
}

type Tasks []*task

func (list Tasks) Len() int {
	return len(list)
}

func (list Tasks) Swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func (list Tasks) Less(a, b int) bool {
	if list[a].CreatedAt == nil {
		return true
	} else if list[b].CreatedAt == nil {
		return false
	}
	return list[a].CreatedAt.Before(*list[b].CreatedAt)
}

func (c *client) Tasks(opts ...func(*TaskOpt)) (out Tasks, err error) {
	i := &TaskOpt{}

	for _, f := range opts {
		f(i)
	}
	path := "tasks"
	if enc := i.Encode(); enc != "" {
		path += "?" + enc
	}
	return out, c.load("GET", path, nil, &out)

}

func (c *client) Task(id int) (t *task, err error) {
	return t, c.load("GET", "tasks/"+strconv.Itoa(id), nil, &t)
}

func (c *client) DeleteTask(id int) error {
	t, err := c.Task(id)
	if err != nil {
		return err
	}
	_, err = c.request("DELETE", "tasks/"+strconv.Itoa(id)+"?revision="+strconv.Itoa(t.Revision), nil)
	return err
}

func (c *client) CompleteTask(id int) error {
	return c.UpdateTask(&UpdateTask{ID: id, Completed: b2p(true)})
}

type UpdateTask struct {
	ID        int     `json:"id,omitempty"`
	Revision  int     `json:"revision,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
	Title     *string `json:"title,omitempty"`
}

func (c *client) UpdateTask(update *UpdateTask) error {
	t, err := c.Task(update.ID)
	if err != nil {
		return err
	}
	update.Revision = t.Revision
	b, err := json.Marshal(update)
	if err != nil {
		return err
	}
	_, err = c.request("PATCH", "tasks/"+strconv.Itoa(update.ID), bytes.NewReader(b))
	return err
}
