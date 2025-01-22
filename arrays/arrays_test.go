package main

import (
	"testing"
)

func TestArraySum(t *testing.T) {
	t.Run("collections of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}

	})

}
