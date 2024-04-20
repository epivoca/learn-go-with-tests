package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of ant size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sums of all given slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertEqualSlices(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sums tails of all given slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertEqualSlices(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T){
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}
		assertEqualSlices(t, got, want)
	})
	t.Run("safely sum empty slice with non empty", func(t *testing.T){
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertEqualSlices(t, got, want)
	})
}

func assertEqualSlices(t testing.TB, got, want []int) {
	t.Helper()

	// Note that this function expects the elements to be comparable. So, it can't be applied to slices
	//  with non-comparable elements like 2D slices
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
