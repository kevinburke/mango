package mango

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

var ua = fmt.Sprintf("github.com/kevinburke/mango/%s go/%s", Version, runtime.Version())

func (c *Client) makeRequest(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	// These both block
	if method == "GET" || method == "HEAD" {
		if c.readWindow != nil {
			if err := c.readWindow.Add(ctx); err != nil {
				return nil, err
			}
		}
	} else {
		if c.writeWindow != nil {
			if err := c.writeWindow.Add(ctx); err != nil {
				return nil, err
			}
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Content-Type", "application/json")
	if c.key != "" {
		req.Header.Set("Authorization", "Key "+c.key)
	}
	return c.client.Do(req)
}
