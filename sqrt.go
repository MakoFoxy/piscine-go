package piscine

func Sqrt(nb int) int {
	a := 1
	for i := 1; i <= nb; i++ {
		a = i * i
		if a == nb {
			return i
		}
	}
	return 0
}
