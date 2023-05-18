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

	readCapacity        int64
	readRemaining       int64
	readFillRate        int64
	readLastFilled      time.Time
	readMu              sync.Mutex
	readCond            *sync.Cond
	readRefillerRunning bool

	writeCapacity        int64
	writeRemaining       int64
	writeFillRate        int64
	writeLastFilled      time.Time
	writeMu              sync.Mutex
	writeCond            *sync.Cond
	writeRefillerRunning bool
}

// Set read capacity in requests per second
func (c *Client) SetReadCapacity(rps int64) {
	if c.readRefillerRunning {
		panic("cannot set read capacity more than once")
	}
	c.readFillRate = rps
	c.readCapacity = rps
	c.readRemaining = rps
	c.readLastFilled = time.Now()
	c.readCond = sync.NewCond(&c.readMu)
	c.readRefillerRunning = true
	go c.readRefill()
}

// Set write capacity in requests per second
func (c *Client) SetWriteCapacity(rps int64) {
	if c.writeRefillerRunning {
		panic("cannot set write capacity more than once")
	}
	c.writeFillRate = rps
	c.writeCapacity = rps
	c.writeRemaining = rps
	c.writeLastFilled = time.Now()
	c.writeCond = sync.NewCond(&c.writeMu)
	c.writeRefillerRunning = true
	go c.writeRefill()
}

func (c *Client) readRefill() {
	for {
		c.readMu.Lock()
		c.readUpdate()
		c.readMu.Unlock()

		c.readCond.Broadcast()
		dur := time.Duration(1000/c.readFillRate) * time.Millisecond
		time.Sleep(dur)
	}
}

func (c *Client) writeRefill() {
	for {
		c.writeMu.Lock()
		c.writeUpdate()
		c.writeMu.Unlock()

		c.writeCond.Broadcast()
		dur := time.Duration(1000/c.writeFillRate) * time.Millisecond
		time.Sleep(dur)
	}
}

// Caller must hold c.readMu
func (c *Client) readUpdate() {
	now := time.Now()
	elapsed := now.Sub(c.readLastFilled).Seconds()
	amount := int64(elapsed * float64(c.readFillRate))

	c.readRemaining += amount
	if c.readRemaining > c.readCapacity {
		c.readRemaining = c.readCapacity
	}

	c.readLastFilled = now
}

// Caller must hold c.writeMu
func (c *Client) writeUpdate() {
	now := time.Now()
	elapsed := now.Sub(c.writeLastFilled).Seconds()
	amount := int64(elapsed * float64(c.writeFillRate))

	c.writeRemaining += amount
	if c.writeRemaining > c.writeCapacity {
		c.writeRemaining = c.writeCapacity
	}

	c.writeLastFilled = now
}

func (c *Client) canConsumeRead() {
	c.readMu.Lock()
	defer c.readMu.Unlock()
	if c.readCapacity == 0 { // never set capacity
		return
	}

	for c.readRemaining <= 0 {
		c.readCond.Wait()
	}

	c.readRemaining--
}

func (c *Client) canConsumeWrite() {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	if c.writeCapacity == 0 { // never set capacity
		return
	}

	for c.writeRemaining <= 0 {
		c.writeCond.Wait()
	}

	c.writeRemaining--
}

// WritesAvailable returns the number of writes available at this instant
func (c *Client) WritesAvailable() int64 {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	return c.writeRemaining
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
