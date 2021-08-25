package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for _, e := range a {
		if f(e) > 0 {
			return false
		} else {
			return true
		}
	}
}
