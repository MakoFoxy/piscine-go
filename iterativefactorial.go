package piscine

func IterativeFactorial(nb int) int {
	if nb > 25 || nb < 0 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		b := 1
		for i := nb; i >= 1; i-- {
			b = b * i
		}
		return b
	}
}
