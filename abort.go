package piscine

func Abort(a, b, c, d, e int) int {
	dev := [5]int{a, b, c, d, e}
	for i := 0; i < len(dev)-1; i++ {
		for i := 0; i < len(dev)-1; i++ {
			if dev[i] > dev[i+1] {
				dev[i], dev[i+1] = dev[i+1], dev[i]
			}
		}
	}
	return dev[2]
}
