package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	for _, n := range arg {
		if n == "01" || n == "galaxy" || n == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
