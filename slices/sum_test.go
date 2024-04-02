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

		// Note that this function expects the elements to be comparable. So, it can't be applied to slices
		//  with non-comparable elements like 2D slices
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
