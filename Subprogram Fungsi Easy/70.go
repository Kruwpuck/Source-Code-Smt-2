// Buatlah fungsi saja
func perkalianDigit(x int) int {
	// fungsi akan mengembalikan perkalian digit-digit bilangan berdigit 3
	var d1, d2, d3 int
	d1 = x / 100
	d2 = x / 10 % 10
	d3 = x % 10
	return d1 * d2 * d3
}