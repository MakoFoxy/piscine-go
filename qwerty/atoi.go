package piscine

func Atoi(s string) int {
	newInt := 0
	var number int
	plussorminus := 1
	for posnum, symbol := range s {
		if posnum == 0 && symbol == 43 {
			continue
		} else if posnum == 0 && symbol == 45 {
			plussorminus = -1
		} else if symbol >= 48 && symbol <= 57 {
			for i := '0'; symbol > i; i++ {
				number++
			}
			newInt = (newInt * 10) + number
			number = 0
		} else {
			return 0
		}
	}
	return newInt * plussorminus
}
