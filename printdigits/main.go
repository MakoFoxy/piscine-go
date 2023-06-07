package main

import "github.com/01-edu/z01"

func main() {
	num := [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, n := range num {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
