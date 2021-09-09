package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Writing a string to the buffer results in the string being saved.", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		AssertSolution(t, got, want)
	})
}

func AssertSolution(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
