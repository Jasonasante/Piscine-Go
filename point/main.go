package main

import "github.com/01-edu/z01"

func printStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
func main() {
	r := "x = 42, y = 21"
	printStr(r)
}
