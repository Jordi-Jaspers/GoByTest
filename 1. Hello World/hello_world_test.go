package main

import "testing"

/*
	Writing a test is just like writing a function, with a few rules
	It needs to be in a file with a name like xxx_test.go
	The test function must start with the word Test
	The test function takes one argument only t *testing.T
	In order to use the *testing.T type, you need to import "testing"
*/
func TestHello(t *testing.T) {
	got := Hello("Jordi")
	want := "Hello Jordi! How are you doing?"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
