package piscine

func LastRune(s string) rune {
	a := []rune(s)

	b := len(a) - 1

	return rune(a[b])
}
