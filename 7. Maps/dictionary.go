package dictionary

import (
	"errors"
	"regexp"
)

// An interesting property of maps is that you can modify them without passing as an address to it
// "A map value is a pointer to a runtime.hmap structure."
// So when you pass a map to a function/method, you are indeed copying it,
// but just the pointer part, not the underlying data structure that contains the data.
// A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading,
// but attempts to write to a nil map will cause a runtime panic.
type Dictionary map[string]string

// ========= EXCEPTIONS ===========
var definitionNotFoundException = "No definition has been found for the given key word"
var badFormatingException = "The key word contains invalid characters"
var wordAlreadyExistException = "The word already exist"

// Check whether the string is alphabetical via regular expression.
var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

// the second variable that is returned in Go maps is a boolean that returns
// true if nothing was found in the defined map. else false.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New(definitionNotFoundException)
	}
	return definition, nil
}

// Adding to a map is also similar to an array.
// You just need to specify a key and set it equal to a value.
func (d Dictionary) Add(word string, definition string) error {
	if isStringAlphabetic(word) {
		_, err := d.Search(word)
		if err.Error() == definitionNotFoundException {
			d[word] = definition
			return nil
		} else {
			return errors.New(wordAlreadyExistException)
		}
	}
	return errors.New(badFormatingException)
}

func (d Dictionary) Update(word string, definition string) error {
	if isStringAlphabetic(word) {
		d[word] = definition
		return nil
	}
	return errors.New(badFormatingException)
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
