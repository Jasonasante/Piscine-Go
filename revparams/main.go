package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []string(os.Args)
	for i := len(args) - 1; i > 0; i-- {
		revlist := args[i]
		if i > 0 {
			for _, char := range revlist {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
	}
}
