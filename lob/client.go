package lob

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Client abstracts interacting with lob.com's api
type Client struct {
	apiKey    string
	netClient *http.Client
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewClient lob.com api client abstraction
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		netClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// IsTest returns true if the client was configured with a test api key
func (c *Client) IsTest() bool {
	return strings.HasPrefix(c.apiKey, "test_")
}

func (c *Client) config(r *http.Request) {
	r.SetBasicAuth(c.apiKey, "")
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")
}
