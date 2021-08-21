package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, arg := range os.Args {
		if i > 0 {
			for _, char := range arg {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
	}
}
