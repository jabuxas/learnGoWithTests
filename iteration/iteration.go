package iteration

// Repeats given character n times
func Repeat(character string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += character
	}
	return
}
