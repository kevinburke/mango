package mango

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := &Client{}
	client.SetReadCapacity(100)
	for i := 0; i < 250; i++ {
		client.canConsumeRead()
		fmt.Printf("Operation %d allowed\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}
