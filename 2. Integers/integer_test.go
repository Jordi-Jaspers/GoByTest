package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("The sum of integers is correctly processed.", func(t *testing.T) {
		sum := add(2, 2)
		expected := 4
		assertCorrectCalculations(t, expected, sum)
	})
}

func assertCorrectCalculations(t testing.TB, sum int, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("The expected sum was '%d' and we got '%d'.", expected, sum)
	}
}
