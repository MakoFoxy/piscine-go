package piscine

func Any(f func(string) bool, a []string) bool {
	for n := range a {
		if f(a[n]) == true {
			return true
		}
	}
	return false
}
