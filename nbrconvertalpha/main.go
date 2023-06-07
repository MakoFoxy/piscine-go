package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var isUpper bool
	arguments := os.Args[1:]
	for ind, argument := range arguments {
		if ind == 0 && argument == "--upper" {
			isUpper = true
			continue
		}

		arg_int := BasicAtoi2(argument)

		var tt rune
		if arg_int == 0 {
			tt = ' '
		} else {
			if isUpper {
				tt = rune(arg_int + 64)
			} else {
				tt = rune(arg_int + 96)
			}
		}
		z01.PrintRune(tt)

	}
	if len(arguments) > 0 {
		z01.PrintRune('\n')
	}
}

func BasicAtoi2(s string) int32 {
	var f int32 = 0
	for i := 0; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			return 0
		} else {
			f = f*10 + (int32(s[i]) - 48)
		}
	}
	if f > 26 {
		return 0
	}
	return f
}
