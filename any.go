package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for _, e := range a {
		result = f(e)
		if result == true {
			return true
		}
	}
	return result
}
