package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("The perimeter of the rectangle can be calculated.", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		assertSolution(t, got, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("The are of a rectangle can be calculated.", func(t *testing.T) {
		got := Area(10.0, 10.0)
		want := 100.0

		assertSolution(t, got, want)
	})
}

func BenchmarkClass(b *testing.B) {
	b.Run("Benchmarking the perimeter functionality", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Perimeter(10.0, 10.0)
		}
	})
	b.Run("Benchmarking the are fucntionality", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Area(10.0, 10.0)
		}
	})
}

func assertSolution(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("the expected value '%f' does not correspond to the computed value '%f'", got, want)
	}

}
