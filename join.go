package piscine

func Join(strs []string, sep string) string {
	var a string
	for i := range strs {
		a = a + strs[i]
		if i != len(strs)-1 {
			a = a + sep
		}
	}
	return a
}
