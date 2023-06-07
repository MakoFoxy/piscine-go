package piscine

func Rot14(s string) string {
	var rot string

	for _, x := range s {
		if x >= 'a' && x <= 'z' {
			if x+14 > 'z' {
				rot += string(x + 14 - 26)
				continue
			}
			rot += string(x + 14)
		} else if x >= 'A' && x <= 'Z' {
			if x+14 > 'Z' {
				rot += string(x + 14 - 26)
				continue
			}
			rot += string(x + 14)

		} else {
			rot += string(x)
		}
	}
	return rot
}
