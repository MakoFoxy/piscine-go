package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, b := range a {
		str := b
		for _, s := range str {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}
