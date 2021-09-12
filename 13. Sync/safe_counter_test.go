package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing 3 times leaves the counter at 3", func(t *testing.T) {
		wantedCount := 3
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		AssertCount(t, counter, wantedCount)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		AssertCount(t, counter, wantedCount)
	})
}

func AssertCount(t testing.TB, counter *Counter, wantedCount int) {
	t.Helper()
	if counter.Value() != wantedCount {
		t.Errorf("got %d, want %d", counter.Value(), wantedCount)
	}
}
