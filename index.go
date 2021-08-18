package piscine

func Index(s string, toFind string) int {
	string := []rune(s)
	substring := []rune(toFind)
	counter := 0
	if len(toFind) > len(s) {
		return -1
	} else {
		for i := 0; i <= len(s)-len(toFind); i++ {
			for j := 0; j < len(toFind); j++ {
				if string[j+i] == substring[j] {
					counter += 1
				} else {
					counter = 0
				}
				if counter == len(toFind) {
					return i
				}
			}
		}
		return -1
	}
}
