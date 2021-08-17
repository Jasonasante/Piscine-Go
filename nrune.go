package piscine

func NRune(s string, n int) rune {
	arune := []rune(s)
	if n < 0 || n > len(s) {
		return 0
	}
	return arune[n-1]
}
