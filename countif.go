package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result bool
	counter := 0
	for _, i := range tab {
		result = f(i)
		if result == true {
			counter += 1
		}
	}
	return counter
}
