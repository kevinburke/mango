package mango

import (
	"net/http"
	"time"
)

// Client represents the main Mango client, used to make requests to the Manifold API
type Client struct {
	client http.Client
	key    string
	url    string

	readWindow  *slidingWindow
	writeWindow *slidingWindow
}

// Set read capacity in requests per second
func (c *Client) SetReadCapacity(maxRequests int, window time.Duration) {
	c.readWindow = newSlidingWindow(maxRequests, window)
}

// Set write capacity in requests per *minute*
func (c *Client) SetWriteCapacity(maxRequests int, window time.Duration) {
	c.writeWindow = newSlidingWindow(maxRequests, window)
}

func (c *Client) ReadsAvailableNow() int {
	return c.readWindow.AvailableNow()
}

func (c *Client) WritesAvailableNow() int {
	return c.writeWindow.AvailableNow()
}

func New(client http.Client, apiURL string, key string) *Client {
	return &Client{
		client: client,
		url:    apiURL,
		key:    key,
	}
}
