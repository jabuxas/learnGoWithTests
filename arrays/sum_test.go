package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertCorrectStatement(t, got, want, numbers)
	})
}

func assertCorrectStatement(tb testing.TB, got, want int, numbers []int) {
	tb.Helper()
	if got != want {
		tb.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}
