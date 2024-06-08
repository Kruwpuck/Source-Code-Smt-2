// Buatlah fungsi saja

func numToTeks(x int) string {
	// Mengembalikan nilai berupa string "negatif", "netral", atau "positif"
	// berdasarkan nilai num, yaitu -1, 0, 1.
	if x > 0 {
		return "positif"
	} else if x < 0 {
		return "negatif"
	} else {
		return "netral"

	}
}