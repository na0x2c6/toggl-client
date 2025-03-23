package client

import (
	"net/http"
)

func (c *Client) setBaseAuth(req *http.Request) {
	req.SetBasicAuth(c.Email, c.Password)
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	c.setBaseAuth(req)

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-type", "application/json; charset=utf-8")

	return c.HTTPClient.Do(req)
}
