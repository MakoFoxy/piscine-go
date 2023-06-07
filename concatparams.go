package piscine

func ConcatParams(args []string) string {
	a := ""
	for i := 0; i < len(args); i++ {
		if i < len(args)-1 {
			a = a + args[i] + "\n"
		} else {
			a = a + args[i]
		}
	}
	return a
}
