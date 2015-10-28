package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const defaultEndpoint = "https://a.wunderlist.com/api/v1"

func loadClient() (*client, error) {
	c := &client{Endpoint: defaultEndpoint}
	t, err := loadTransport()
	if err != nil {
		return nil, err
	}
	t.client = c
	c.authClient = &http.Client{Transport: t}
	return c, nil
}

type client struct {
	Endpoint   string
	Client     *http.Client // use this to hook in your custom client
	authClient *http.Client
}

func (c *client) load(method, path string, payload io.Reader, i interface{}) error {
	rsp, err := c.request(method, path, payload)
	if err != nil {
		return err
	}
	return json.NewDecoder(rsp.Body).Decode(&i)
}

func (c *client) request(method, path string, payload io.Reader) (*http.Response, error) {
	u := c.Endpoint + "/" + path
	dbg.Printf("[REQ] method=%s url=%s", method, u)
	req, err := http.NewRequest(method, u, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	rsp, err := c.authClient.Do(req)
	if err != nil {
		return nil, err
	}
	if rsp.Status[0] != '2' {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("got status %s but expected 2x. body=%s", rsp.Status, string(b))
	}
	return rsp, nil
}
