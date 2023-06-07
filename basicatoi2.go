package piscine

func BasicAtoi2(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			c = c*10 + int(s[i]-48)
		} else {
			return 0
		}
	}
	return c
}
