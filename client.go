package mango

import (
	"net/http"
	"sync"
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

var lock = &sync.Mutex{}
var mcInstance *Client // TODO: figure out whether this should really be a singleton or not
