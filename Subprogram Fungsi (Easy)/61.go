// Buatlah fungsi saja

func fungsiY(a float64) float64 {
	// fungsi akan mengembalikan nilai y = 4x^3 + 3x^2 + 2x + 1 jika diberikan x
	var b float64
	b = 1 + (4 * (a * a * a)) + (3 * (a * a)) + (2 * (a * a))
	return b
}