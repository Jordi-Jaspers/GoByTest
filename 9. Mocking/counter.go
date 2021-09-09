package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// ============ Defining and implementing the interface ================

// Creating an interface with defined method it needs.
type Timer interface {
	Sleep()
}

// Implementing the interface by creating a struct and assigning the same method as defined in the struct.
// Inserting the a function as a variable in the struct. So, when we pass 'time.Sleep' to this configurable sleeper.
// It will execute this function.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (t *ConfigurableSleeper) Sleep() {
	t.sleep(t.duration)
}

// ============ Code ================

const COUNTER_START = 3
const FINAL_WORD = "Go!"

func Countdown(writer io.Writer, timer Timer) {

	for i := COUNTER_START; i > 0; i-- {
		fmt.Fprintln(writer, i)
		timer.Sleep()
	}
	fmt.Fprint(writer, FINAL_WORD)
}

func main() {
	// Cannot initialize the interface, so we intialize the implementation of the interface.
	timer := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, timer)
}

// ============== NOTES =====================
/*
Normally a lot of mocking points to bad abstraction in your code. Usually it is a sign of:

1. The thing you are testing is having to do too many things (because it has too many dependencies to mock)
	- Break the module apart so it does less

2. Its dependencies are too fine-grained
	- Think about how you can consolidate some of these dependencies into one meaningful module

3. Your test is too concerned with implementation details
	- Favour testing expected behaviour rather than the implementation

General Rules of Thumb:
	1. 	Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behaviour.
		Test the public behaviour. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.

	2.	Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means a tighter coupling between your test code and the implementation.
		Be sure you actually care about these details if you're going to spy on them

	3. I feel like if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design.
*/
