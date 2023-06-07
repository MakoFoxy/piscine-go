// Given a slice of integers, return indexes of the two numbers such that they add up to a specific target.

// If there are more than one solution, return the first one.

// If there are no solutions, return nil.

package piscine

func Twosum(a []int, b int) []int {
	var arr []int
	count := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if j == i {
				continue
			} else {
				if a[i]+a[j] == b {
					count += 1
					if count == 1 {
						arr = append(arr, i)
						arr = append(arr, j)
					}
					break
				}
			}
		}
	}
	return arr
}
