package wlclient

import "strconv"

type TaskPosition struct {
	ID       int    `json:"id,omitempty"`
	Values   []int  `json:"values,omitempty"`
	Revision int    `json:"revision,omitempty"`
	ListID   int    `json:"list_id,omitempty"`
	Type     string `json:"type,omitempty"`
}

func (c *Client) TaskPositions(listID int) (out []*TaskPosition, err error) {
	return out, c.load("GET", "/task_positions?list_id="+strconv.Itoa(listID), nil, &out)
}
