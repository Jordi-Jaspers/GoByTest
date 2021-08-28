package interfaces

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-1

var rectangle = Rectangle{10.0, 10.0}
var circle = Circle{10.0}

func TestPerimeter(t *testing.T) {
	t.Run("The perimeter of the rectangle can be calculated.", func(t *testing.T) {
		want := 40.0
		assertSolutionPerimeter(t, rectangle, want)
	})
	t.Run("The perimeter of the circle can be calculated.", func(t *testing.T) {
		want := 62.8
		assertSolutionPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("The area of a rectangle can be calculated.", func(t *testing.T) {
		want := 100.0
		assertSolutionArea(t, rectangle, want)
	})
	t.Run("The area of the circle can be calculated.", func(t *testing.T) {
		want := 314.2
		assertSolutionArea(t, circle, want)
	})
}

func BenchmarkClass(b *testing.B) {
	b.Run("Benchmarking the perimeter functionality", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rectangle.Perimeter()
			circle.Perimeter()
		}
	})
	b.Run("Benchmarking the are fucntionality", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rectangle.Area()
			circle.Area()
		}
	})
}

func assertSolutionArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if math.Abs(got-want) >= float64EqualityThreshold {
		t.Errorf("the expected value '%.1f' does not correspond to the computed value '%.1f'", got, want)
	}
}

func assertSolutionPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()
	if math.Abs(got-want) >= float64EqualityThreshold {
		t.Errorf("the expected value '%.1f' does not correspond to the computed value '%.1f'", got, want)
	}
}
