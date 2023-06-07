package piscine

func Swap(a *int, b *int) {
	A := *a
	B := *b

	*b = A
	*a = B
}
