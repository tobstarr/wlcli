package wlclient

import (
	"fmt"
	"strconv"
	"time"
)

type List struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Title     string    `json:"title,omitempty"`
	ListType  string    `json:"list_type,omitempty"`
	Type      string    `json:"type,omitempty"`
	Revision  int       `json:"revision,omitempty"`
}

func (c *Client) Inbox() (*List, error) {
	lists, err := c.Lists()
	if err != nil {
		return nil, err
	}
	for _, l := range lists {
		if l.ListType == "inbox" {
			return l, nil
		}
	}
	return nil, fmt.Errorf("inbox not found")
}

func (c *Client) Lists() (out []*List, err error) {
	return out, c.load("GET", "lists", nil, &out)
}

func (c *Client) List(id int) (l *List, err error) {
	return l, c.load("GET", "lists/"+strconv.Itoa(id), nil, &l)

}
