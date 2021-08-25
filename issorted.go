package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for e := range a {
		for k := 0; k == len(a); k++ {
			if f(a[e], a[e+1]) > 0 {
				return false
			}
		}
	}
	return true
}
