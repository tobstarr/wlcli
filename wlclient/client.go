package wlclient

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const defaultEndpoint = "https://a.wunderlist.com/api/v1"

// Documentation at: https://developer.wunderlist.com/documentation

func New(clientID, accessToken string) (*Client, error) {
	if clientID == "" || accessToken == "" {
		return nil, fmt.Errorf("clientID=%q and accessToken=%q must bost be present", clientID, accessToken)
	}
	c := &Client{Endpoint: defaultEndpoint}
	t := &transport{ClientID: clientID, AccessToken: accessToken, client: c}
	t.client = c
	c.authClient = &http.Client{Transport: t}
	return c, nil

}

type Client struct {
	Endpoint   string
	Client     *http.Client // use this to hook in your custom client
	authClient *http.Client
}

func (c *Client) load(method, path string, payload io.Reader, i interface{}) error {
	rsp, err := c.request(method, path, payload)
	if err != nil {
		return err
	}
	return json.NewDecoder(rsp.Body).Decode(&i)
}

func (c *Client) request(method, path string, payload io.Reader) (*http.Response, error) {
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
