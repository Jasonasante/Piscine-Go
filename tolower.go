package piscine

func ToLower(s string) string {
	astr := []rune(s)

	for i, j := range astr {
		if j >= 65 && j <= 90 {
			astr[i] = astr[i] + 32
		}
	}
	return string(astr)
}
