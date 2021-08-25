package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, e := range a {
		result = append(result, f(e))
	}
	return result
}
