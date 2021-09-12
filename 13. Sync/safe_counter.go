package main

import "sync"

//We want to make a counter which is safe to use concurrently.
// We'll start with an unsafe counter and verify its behaviour works in a single-threaded environment

// Mutex is a mutual exclusion that will let you lock a certain property until it is unlocked again. This helps with concurrency.
// whenever a go routine would like to access the object at the same time it is impossible unless the object gets unlocks first by
// another goroutine that is currently accessing the object.
type Counter struct {
	sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.Lock()
	c.value++
	c.Unlock()
}
