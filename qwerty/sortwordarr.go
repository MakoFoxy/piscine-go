// write  a function SortWordArr that sorts by ascii (in ascending order) a string array.

package piscine

func SortWordArr(a []string) {
	for {
		swapped := false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
