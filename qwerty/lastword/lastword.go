// Write a program that takes a string and displays its last word, followed by a newline ('\n').

//  A word is a section of string delimited by spaces or by the start/end of the string.

//  The output will be followed by a newline ('\n').

//  If the number of parameters is different from 1, or if there are no word, the program displays a newline ('\n').

package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	arg := os.Args[1:]
// 	if len(os.Args) == 0 {
// 		return
// 	}
// 	var revargs []string
// 	for i := len(arg) - 1; i > -1; i-- {
// 		revargs = append(revargs, arg[i])
// 	}
// 	for _, b := range revargs[0] {
// 		z01.PrintRune(b)
// 	}

// 	z01.PrintRune('\n')
// }
func main() {
	if len(os.Args[1:]) == 1 {
		arg := []rune(os.Args[1])
		count := 0
		for i, k := range arg {
			if i < len(arg)-2 && k == 32 {
				count = i + 1
			}
		}
		for i := count; i < len(arg); i++ {
			if arg[i] != 32 {
				z01.PrintRune(arg[i])
			}
		}

	}
}
