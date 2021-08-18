package piscine

func IsPrintable(s string) bool {
	astr := []rune(s)

	for _, i := range astr {
		if i < 21 || i > 126 {
			return false
		}
	}
	return true
}
