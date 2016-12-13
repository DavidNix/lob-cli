package lob

import (
	"net/http"
	"time"
)

// Client abstracts interating with lob.com's api
type Client struct {
	apiKey    string
	netClient *http.Client
}

//NewClient lob.com api client abstraction
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		netClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) config(r *http.Request) {
	r.SetBasicAuth(c.apiKey, "")
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")
}
