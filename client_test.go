package mango

import (
	"testing"
)

func TestClient(t *testing.T) {
	client := &Client{}
	client.SetWriteCapacity(6)
	for i := 0; i < 250; i++ {
		go func() {
			client.canConsumeWrite()
		}()
	}
	if w := client.WritesAvailable(); w > 0 {
		t.Errorf("should not have had writes available but got %q", w)
	}
}
