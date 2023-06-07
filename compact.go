package piscine

func Compact(ptr *[]string) int {
	var a []string

	size := 0
	for _, n := range *ptr {
		if n != "" {
			size++
			a = append(a, n)
		}
	}

	*ptr = a

	return size
}
