package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	// row := y
	// col := x
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (col == 1 && row == 1) || (col == x && row == y && row != 1 && col != 1) {
				z01.PrintRune('A')
			} else if (col == x && row == 1) || (col == 1 && row == y) {
				z01.PrintRune('C')
			} else if (row == 1) || (row == y) || (col == 1) || (col == x) {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
