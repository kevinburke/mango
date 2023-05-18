package mango

import (
	"context"
	"sync"
	"time"
)

type slidingWindow struct {
	mu          sync.Mutex
	maxRequests int
	windowSize  time.Duration
	timestamps  []time.Time
}

func newSlidingWindow(maxRequests int, windowSize time.Duration) *slidingWindow {
	return &slidingWindow{
		maxRequests: maxRequests,
		windowSize:  windowSize,
		timestamps:  make([]time.Time, 0),
	}
}

func (s *slidingWindow) AvailableNow() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	startWindow := now.Add(-s.windowSize)
	newTimestamps := make([]time.Time, 0)

	for _, ts := range s.timestamps {
		if ts.After(startWindow) {
			newTimestamps = append(newTimestamps, ts)
		}
	}

	if len(newTimestamps) < s.maxRequests {
		return s.maxRequests - len(newTimestamps)
	}
	return 0
}

func (s *slidingWindow) Add(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	startWindow := now.Add(-s.windowSize)
	newTimestamps := make([]time.Time, 0)

	for _, ts := range s.timestamps {
		if ts.After(startWindow) {
			newTimestamps = append(newTimestamps, ts)
		}
	}

	if len(newTimestamps) < s.maxRequests {
		newTimestamps = append(newTimestamps, now)
		s.timestamps = newTimestamps
		return nil
	}

	oldestEntry := newTimestamps[0]
	select {
	case <-time.After(time.Until(oldestEntry.Add(s.windowSize))):
		// TODO, not great to have two different entries here, but we hold the
		// lock so it should be okay-ish
		newTimestamps = append(newTimestamps, now)
		s.timestamps = newTimestamps
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
