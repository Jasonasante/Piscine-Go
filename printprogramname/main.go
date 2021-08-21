package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	name := arguments[0]
	runes := []rune(name)
	for i, j := range runes {
		if i > 1 {
			z01.PrintRune(rune(j))
		}
	}
	z01.PrintRune('\n')
}
