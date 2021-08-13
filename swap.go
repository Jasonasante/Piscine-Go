package piscine

func Swap(a *int, b *int) {
	a = *b
	b = *a - 1
}
