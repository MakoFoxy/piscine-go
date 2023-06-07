// Write a function that counts the characters of a string and that returns that count.

package piscine

func StrLen(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}
