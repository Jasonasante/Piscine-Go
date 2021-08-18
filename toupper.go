package piscine

func ToUpper(s string) string {
	astr := []rune(s)

	for i, j := range astr {
		if j >= 97 && j <= 122 {
			astr[i] = astr[i] - 32
		}
	}
	return string(astr)
}
