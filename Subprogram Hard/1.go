func isPrime(bilangan int) bool {
	// Bilangan 1 bukan prima
	if bilangan == 1 {
		return false
	}
	// Cek apakah bilangan habis dibagi oleh bilangan lain selain 1 dan dirinya sendiri
	for i := 2; i*i <= bilangan; i++ {
		if bilangan%i == 0 {
			return false
		}
	}
	return true
}