package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	count := 0
	dev := false
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			count++
		}
	}
	if count == 0 || count == len(a)-1 {
		dev = true
	}
	return dev
}
