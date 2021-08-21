package piscine

func MakeRange(min, max int) []int {
	if max-min > 0 {
		result := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			result[i] = i + min
		}
		return result
	} else {
		return nil
	}
}
