package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	a := make([]int, max-min)
	for i := 0; i < (max - min); i++ {
		a[i] = min + i
	}
	return a
}
