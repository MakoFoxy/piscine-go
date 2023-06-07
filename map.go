package piscine

func Map(f func(int) bool, a []int) []bool {
	var rav []bool

	for n := range a {
		if f(a[n]) == true {
			rav = append(rav, true)
		} else {
			rav = append(rav, false)
		}
	}
	return rav
}
