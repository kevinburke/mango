package mango

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/viper"
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

// ClientInstance creates a singleton of the Mango Client.
// It optionally takes a http.Client, base URL, and API key.
//
// If you don't specify a base URL, the default Manifold Markets domain will be used.
//
// If no API key is provided then you will need to specify a `MANIFOLD_API_KEY` in your .env file
//
// Just because you *can* specify an API key here doesn't mean that you *should*!
// Please don't put your API key in code.
func ClientInstance(client *http.Client, url, ak *string) *Client {
	if mcInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if mcInstance == nil {
			if client == nil {
				client = &http.Client{
					Timeout: time.Second * 10,
				}
			}

			if url == nil {
				u := base
				url = &u
			}

			if ak == nil {
				a := apiKey()
				ak = &a
			}

			mcInstance = &Client{
				client: *client,
				key:    *ak,
				url:    *url,
			}
		}
	}
	return mcInstance
}

// DefaultClientInstance returns a singleton of the Mango Client using all default values.
//
// It will use a default http.Client, the primary Manifold domain as the base URL, and
// the value of `MANIFOLD_API_KEY` in your .env file as the API key.
func DefaultClientInstance() *Client {
	return ClientInstance(nil, nil, nil)
}

// Destroy destroys the current singleton of the Mango client.
//
// Useful for testing.
func (mc *Client) Destroy() {
	if mcInstance != nil {
		lock.Lock()
		defer lock.Unlock()
		mcInstance = nil
	}
}

func apiKey() string {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Errorf("fatal error config file: %w", err)
	}

	viper.SetEnvPrefix("MANIFOLD")
	viper.AutomaticEnv()

	return viper.GetString("MANIFOLD_API_KEY")
}
