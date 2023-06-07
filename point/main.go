package main

import "github.com/01-edu/z01"

func main() {
	a := ("x = 42, y = 21")
	for _, b := range a {
		z01.PrintRune(b)
	}
	z01.PrintRune('\n')
}
