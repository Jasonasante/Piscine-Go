package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "quest8.txt"
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("the mistake is: %v\n", err.Error())
		return
	}
	fmt.Printf(string(content))
}
