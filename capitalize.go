package piscine

func Capitalize(s string) string {
	var cap []rune
	l := ToLower(s)
	r := []rune(l)
	for i := 0; i < len(l)-1; i++ {
		if r[0] >= 'a' && r[0] <= 'z' {
			r[0] = r[0] - 32
		}
		if r[i] < 48 || r[i] > 57 && r[i] < 65 || r[i] > 90 && r[i] < 97 || r[i] > 122 {
			if r[i+1] >= 'a' && r[i+1] <= 'z' {
				r[i+1] = r[i+1] - 32
			}
		}
		cap = append(cap, r[i])
	}
	cap = append(cap, r[(len(r)-1)])
	return string(cap)
}
