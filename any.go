package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for _, e := range a {
		f(e)
		result = f(e)
		if result {
			return false
		} else {
			return true
		}
	}
	return result
}
