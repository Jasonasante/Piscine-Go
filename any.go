package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for _, e := range a {
		f(e)
		result = f(e)
	}
	return result
}
