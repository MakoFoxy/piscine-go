package piscine

func AlphaCount(s string) int {
	a := 0
	for _, b := range s {
		if (b >= 97 && b <= 122) || (b >= 65 && b <= 90) {
			a++
		}
	}
	return a
}
