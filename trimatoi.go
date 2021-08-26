package piscine

func TrimAtoi(s string) int {
	myRunes := []rune(s)
	num := 0
	counter := 0
	for _, b := range myRunes {
		if b >= 48 && b <= 57 {
			num = num*10 + int(b-'0')
		}
	}
	for _, v := range myRunes {
		if v >= 48 && v <= 57 {
			counter++
		}
		if v == '-' && counter == 0 {
			num = -(num)
		}
	}
	return num
}
