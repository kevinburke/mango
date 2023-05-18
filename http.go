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
		c.canConsumeRead()
	} else {
		c.canConsumeWrite()
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
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
