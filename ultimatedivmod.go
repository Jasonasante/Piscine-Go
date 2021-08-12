package piscine

func UltimateDivMod(a *int, b *int) {
	var division *int = &a / &b
	var modulo *int = &a % &b
	*a = division
	*b = modulo
}
