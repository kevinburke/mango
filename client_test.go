package mango

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := &Client{}
	client.SetWriteCapacity(100, time.Second)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	errs := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := client.writeWindow.Add(ctx); err != nil {
				errs++
			}
		}()
	}
	wg.Wait()
	if errs < 880 || errs > 920 {
		t.Errorf("expected errs to be about 900, got %d", errs)
	}
}
