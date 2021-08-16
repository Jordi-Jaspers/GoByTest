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

	//t.run runs function f as a subtest of t
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jordi", "")
		want := "Hello Jordi! How are you doing?"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello world' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world! How are you doing?"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour Elodie! Cava?"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English", func(t *testing.T) {
		got := Hello("Elodie", "English")
		want := "Hello Elodie! How are you doing?"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
