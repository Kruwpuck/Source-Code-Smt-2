func cetakBilangan(x, y, a, b int) {
	for i := x; i <= y; i++ {
		if i < 100 || i > 999 { // Pastikan hanya memproses bilangan 3 digit
			continue
		}

		// Ambil digit pertama dan digit terakhir
		digitPertama := i / 100
		digitTerakhir := i % 10

		if digitPertama == a || digitTerakhir == b {
			continue // Abaikan bilangan yang dimulai dengan a atau diakhiri dengan b
		}
		fmt.Println(i)
	}
}
