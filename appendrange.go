package piscine

func AppendRange(min, max int) []int {
	var a []int
	for i := min; i <= max-1; i++ {
		a = append(a, i)
	}
	return a
}
