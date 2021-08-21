package piscine

func ConcatParams(args []string) string {
	a := 0
	for i := range args {
		a = i + 1
	}
	result := make([]string, a)
	for j := 0; j < (a - 1); j++ {
		result[1] = result[1] + args[j] + "\n"
	}
	result[1] = result[1] + args[(a-1)]
	return result[1]
}
