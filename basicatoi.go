package piscine

func BasicAtoi(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		c = c*10 + int(s[i]-48)
	}
	return c
}
