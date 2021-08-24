package piscine

func MakeRange(min, max int) []int {
	if max-min < 0 {
		return nil
	} else {
		result := make([]int, max-min)
		for i := min; i < max; i++ {
			result[i-min] = i
		}
		return result
	}
}
