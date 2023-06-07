package piscine

func Unmatch(a []int) int {
	for _, n := range a {
		i := 0
		for _, b := range a {
			if b == n {
				i++
			}
		}
		if i == 1 || i%2 == 1 {
			return n
		}
	}
	return -1
}
