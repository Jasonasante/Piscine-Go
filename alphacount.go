package piscine

func AlphaCount(s string) int {
	arune := []rune(s)
	counter := 0

	for i := range arune {
		if arune[i] >= 65 && arune[i] <= 90 || arune[i] >= 97 && arune[i] <= 122 {
			counter = counter + 1
		}
	}
	return counter
}
