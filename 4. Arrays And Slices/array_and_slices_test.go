package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Array: collection of specified numbers are added.", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertSum(t, got, want, numbers)
	})
	t.Run("Slice: Collection of undefined size is added.", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := Sum(numbers)
		want := 55

		assertSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Varying numbers of slices can be added.", func(t *testing.T) {
		arr1 := []int{1, 2}
		arr2 := []int{0, 4}

		got := SumAll(arr1, arr2)
		want := []int{3, 4}

		assertSumAll(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("The tail of multiple arrays can be added.", func(t *testing.T) {
		arr1 := []int{1, 2}
		arr2 := []int{0, 4}
		arr3 := []int{3, 4, 6, 5}

		got := SumAllTails(arr1, arr2, arr3)
		want := []int{2, 4, 15}

		assertSumAll(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		arr1 := []int{3, 4, 5}
		arr2 := []int{}

		got := SumAllTails(arr1, arr2)
		want := []int{9, 0}

		assertSumAll(t, got, want)
	})
}

// %v prints the default format of an object.
func assertSum(t testing.TB, got int, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// In general DeepEqual is a recursive relaxation of Go's == operator.
// As DeepEqual traverses the data values it may find a cycle. The
// second and subsequent times that DeepEqual compares two pointer
// values that have been compared before, it treats the values as
// equal rather than examining the values to which they point.
// This ensures that DeepEqual terminates.
func assertSumAll(t testing.TB, got []int, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
