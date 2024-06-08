// Buatlah fungsi saja

func kelulusan(n float64) string {
	var lulus bool
	lulus = n > 40.00
	if lulus {
		return "lulus"
	} else {
		return "tidak lulus"
	}
}