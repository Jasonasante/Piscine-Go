package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) == 1 {
		return true
	}
	if f(a[0], a[1]) >= 0 {
		for k := 0; k < len(a); k++ {
			if f(a[k], a[k+1]) < 0 {
				return false
			}
		}
	}
	if f(a[0], a[1]) <= 0 {
		for j := 0; j < len(a); j++ {
			if f(a[j], a[j+1]) > 0 {
				return false
			}
		}
	}
	return true
}
