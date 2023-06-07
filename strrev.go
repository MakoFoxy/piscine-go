package piscine

func StrRev(s string) string {
	c := ""
	for _, x := range s {
		c = string(x) + c
	}
	return c
}
