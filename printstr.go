package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	schangeable := []byte(s)
	for i := range schangeable {
		z01.PrintRune(rune(i))
	}
}
