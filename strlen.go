package piscine

func StrLen(s string) int {
	astr := []rune(s)
	a := 1
	for index := range astr {
		a = index
	}
	return a + 1
}
