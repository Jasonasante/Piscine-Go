package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	schangeable := []byte(s)
	r := schangeable
	z01.PrintRune(rune(r))
}
