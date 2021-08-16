package piscine

func Sqrt(nb int) int {
	i := 1
	for ; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
