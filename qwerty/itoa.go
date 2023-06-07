// Write a function that simulates the behaviour of the Itoa function in Go. Itoa transforms a number represented as anint in a number represented as a string.
// For this exercise the handling of the signs + or - does have to be taken into account.

package piscine

func Itoa(n int) string {
	var arr []rune
	s := ""
	if n < 0 {
		for n != 0 {
			arr = append(arr, rune(-(n%10)+48))
			n /= 10
		}
	} else if n > 0 {
		for n != 0 {
			arr = append(arr, rune(-(n%10)+48))
			n /= 10
		}
	} else {
		arr = append(arr, '0')
	}
	for i := len(arr) - 1; i >= 0; i-- {
		s += string(arr[i])
	}
	return s
}
