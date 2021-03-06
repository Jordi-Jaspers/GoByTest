package iteration

import "testing"

// testing.T gives you access to test code.
func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	assert(t, repeated, expected)
}

// testing.B gives you access to the cryptically named b.N
// When the benchmark code is executed, it runs b.N times and measures how long it takes
// go test -bench=. will run the test 10 000 000 times
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func assert(t testing.TB, repeated string, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
