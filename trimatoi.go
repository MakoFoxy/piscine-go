package piscine

func TrimAtoi(s string) int {
	b := false
	count := 0
	for _, a := range s {
		if a == '-' && count == 0 {
			b = true
		}
		if a >= '0' && a <= '9' {
			count = count*10 + int(a-48)
		}
	}
	if b {
		return -count
	} else {
		return count
	}
}
