package main

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// Go is really optimized for concurrency which is really helpfull
// An operation that does not block in Go will run in a separate process called a goroutine.
// Think of a process as reading down the page of Go code from top to bottom,
// going 'inside' each function when it gets called to read what it does.
// When a separate process starts it's like another reader begins reading inside the function,
// leaving the original reader to carry on going down the page.
//
// Anonymous functions have a number of features which make them useful,
// two of which we're using above. Firstly, they can be executed at the same time that they're
// declared - this is what the () at the end of the anonymous function is doing.
// Secondly they maintain access to the lexical scope they are defined in -
// all the variables that are available at the point when you declare the anonymous function are
// also available in the body of the function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// To prevent racing problems(writing/reading the same object at the same time)
	// Writing the results to a channel and later reading all the results from this channel.
	// Finally add all these results to the original map and return this.
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
