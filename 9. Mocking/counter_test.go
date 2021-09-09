package main

import (
	"bytes"
	"reflect"
	"testing"
)

// Defining a new implementation of the interface we created by implementing the same methods.
// Only the Sleep() method will execute a counter instead of a sleep.
// In other words mocking the interface we used.

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

// io.writer interface has a Write method we would like to Mock.
/*
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
*/
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

// So Mocking a call in Golang is creating a struct that implements the same Method that is called.

func TestCountDown(t *testing.T) {
	t.Run("Counter correctly counts down from 3 to 1 with a 1 second delay", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spyOperations := &SpyCountdownOperations{}

		Countdown(spyOperations, spyOperations)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spyOperations.Calls) {
			t.Errorf("wanted calls %v got %v", want, spyOperations.Calls)
		}
	})
}
