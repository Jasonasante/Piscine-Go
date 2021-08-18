package piscine

func IsUpper(s string) bool {
	astr := []rune(s)

	for _, i := range astr {
		if i < 65 || i > 90 {
			return false
		}
	}
	return true
}
