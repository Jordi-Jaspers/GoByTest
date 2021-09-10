package main

import (
	"errors"
	"net/http"
	"time"
)

const TIMEOUT = 10 * time.Second

var noResponseException = "There was no response from any of the different URLs"

func Racer(url1 string, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, TIMEOUT)
}

func ConfigurableRacer(url1 string, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", errors.New(noResponseException)
	}
}

// If you recall from the concurrency chapter, you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.
// What select lets you do is wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
// We use ping in our select to set up two channels for each of our URLs. Whichever one writes to its channel first will have its code executed in the select, which results in its URL being returned (and being the winner).
// After these changes, the intent behind our code is very clear and the implementation is actually simpler.
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
