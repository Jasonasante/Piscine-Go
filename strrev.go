package piscine

func StrRev(s string) string {
	var revstr []rune
	astr := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		revstr = append(revstr, astr[i])
	}
	return string(revstr)
}
