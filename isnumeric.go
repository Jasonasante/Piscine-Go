package piscine

func IsNumeric(s string) bool {
	astr := []rune(s)

	for _, i := range astr {
		if i < 48 || i > 57 {
			return false
		}
	}
	return true
}
