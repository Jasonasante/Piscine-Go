package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	schangeable := []byte(s)
	for index, i := range schangeable {
		z01.PrintRune(rune(s))
	}
}
