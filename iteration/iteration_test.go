package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		iDontWannaIf(t, repeated, expected)
	})
	t.Run("banana test", func(t *testing.T) {
		repeated := "ba" + Repeat("na", 2)
		expected := "banana"
		iDontWannaIf(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeated := Repeat("k", 35)
	fmt.Println(repeated)
	// Output: kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func iDontWannaIf(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
