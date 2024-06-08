func perkalianLoop(n, m int) int {
	var hasil int = 0
	var absM int

	// Menghitung nilai absolut dari m
	if m < 0 {
		absM = -m
	} else {
		absM = m
	}

	for i := 0; i < absM; i++ {
		hasil += n
	}

	// Jika m negatif, hasilnya juga harus negatif
	if m < 0 {
		hasil = -hasil
	}

	return hasil
}
