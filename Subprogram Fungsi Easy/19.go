// Buatlah fungsi saja

func fx(x float64) float64 {
	var hasil float64
	hasil = (((x * x * x) + (3.0 * x)) / (((x * x * x * x) - (3.0 * (x * x))) + 4.0))
	return hasil
}