package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	count := 0
	for range args {
		count++
	}
	if count == 0 {
		fmt.Println("File name missing")
		return
	} else if count > 1 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := args[0]
	content, err := os.Open(fileName)
	if err != nil {
		fmt.Println()
		return
	}
	arr := make([]byte, 14)
	content.Read(arr)
	fmt.Println(string(arr))
}
