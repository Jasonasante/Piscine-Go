package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := "quest8.txt"
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("the mistke is: %v\n", err.Error())
		return
	}
	fmt.Println(string(content))
}