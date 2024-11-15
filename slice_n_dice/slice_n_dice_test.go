package slice_n_dice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sums collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got '%d' want '%d', given '%v", got, want, numbers)
		}
	})
	t.Run("sums collections of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' want '%d', given '%v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum all given slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}

func TestSumTails(t *testing.T) {
	t.Run("sums all values in a slice after the first and returns a slice of the totals as expected", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4, 5})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
	t.Run("handles empty slice gracefully", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}
