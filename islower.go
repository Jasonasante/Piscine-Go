package piscine

func IsLower(s string) bool {
	astr := []rune(s)

	for _, i := range astr {
		if i < 97 || i > 122 {
			return false
		}
	}
	return true
}
