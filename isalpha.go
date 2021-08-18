package piscine

func IsAlpha(s string) bool {
	astr := []rune(s)

	for _, i := range astr {
		if i < 30 || i > 57 && i < 65 || i > 90 && i < 97 || i > 122 {
			return false
		}
	}
	return true
}
