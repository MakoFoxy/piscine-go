package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	// row := y
	// col := x
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (col == 1 && row == 1) || (col == x && row == y && col != 1 && row != 1) {
				z01.PrintRune('/')
			} else if (col == 1 && row == y) || (col == x) && (row == 1) {
				z01.PrintRune(92)
			} else if col == 1 || row == 1 || col == x || row == y {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
