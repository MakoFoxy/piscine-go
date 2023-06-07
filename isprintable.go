package piscine

func checkprint(a rune) bool {
	if a >= 0 && a <= 31 {
		return true
	}
	return false
}

func IsPrintable(str string) bool {
	r := []rune(str)

	for i := range r {
		if checkprint(r[i]) == true {
			return false
		}
	}
	return true
}
