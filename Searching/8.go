func searchVokal(x string) bool {
	// Iterasi setiap karakter dalam string
	for i := 0; i < len(x); i++ {
		char := x[i]
		// Periksa huruf vokal dalam huruf kecil dan besar
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
			char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			return true
		}
	}
	return false
}