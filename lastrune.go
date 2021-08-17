package piscine

func LastRune(s string) rune {
	arune := []rune(s)
	return arune[len(s)-1]
}
