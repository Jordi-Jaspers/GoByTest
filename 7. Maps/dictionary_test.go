package dictionary

import "testing"

const MAP_VALUE = "the value of the key test"
const MAP_KEY = "test"

var dictionary = Dictionary{MAP_KEY: MAP_VALUE}

func TestSearch(t *testing.T) {
	t.Run("Correctly search a word in a dictionary", func(t *testing.T) {
		got, err := dictionary.Search(MAP_KEY)

		AssertNonError(t, err)
		AssertLookup(t, got, MAP_VALUE, err)
	})
	t.Run("Unkown words cannot be found and notifies the user", func(t *testing.T) {
		_, err := dictionary.Search("unfound")

		AssertError(t, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("A new word can be added to the library", func(t *testing.T) {
		var newWord = "fabulous"
		var description = "The creator of this code"

		_ = dictionary.Add(newWord, description)
		got, err := dictionary.Search(newWord)

		AssertNonError(t, err)
		AssertLookup(t, got, description, err)
	})
	t.Run("A new word with an illegal character returns an error", func(t *testing.T) {
		var newWord = "wrong kind of fabulous!"
		var description = "This ain't gon' compile buddy"

		err := dictionary.Add(newWord, description)

		AssertBadFormattingError(t, err)
	})
	t.Run("An existing word cannot be added multiple times or the definition cannot be changed", func(t *testing.T) {
		var newWord = "fabulous"
		var description1 = "The creator of this code"
		var description2 = "Don't update it with this"

		_ = dictionary.Add(newWord, description1)
		got, err := dictionary.Search(newWord)

		AssertNonError(t, err)
		AssertLookup(t, got, description1, err)

		_ = dictionary.Add(newWord, description2)
		got, err = dictionary.Search(newWord)

		AssertNonError(t, err)
		AssertLookup(t, got, description1, err)
	})
}

func AssertBadFormattingError(t testing.TB, err error) {
	if err == nil {
		t.Fatal("Expected to receive an error")
	}
	if err.Error() != badFormatingException {
		t.Errorf("got %q, want %q \n", err.Error(), badFormatingException)
	}
}

func AssertError(t testing.TB, err error) {
	if err == nil {
		t.Fatal("Expected to receive an error")
	}
	if err.Error() != definitionNotFoundException {
		t.Errorf("got %q, want %q \n", err.Error(), definitionNotFoundException)
	}
}

func AssertNonError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func AssertLookup(t testing.TB, got string, want string, err error) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
