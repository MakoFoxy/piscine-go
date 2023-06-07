package piscine

func ForEach(f func(int), a []int) {
	for s := range a {
		f(a[s])
	}
}
