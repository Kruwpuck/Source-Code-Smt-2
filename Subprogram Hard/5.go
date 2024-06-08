func uangKakak(tabPertama, n int) int {
	/*
		I.S. terdefinisi dua buah bilangan bulat yang menyatakan tabungan di minggu pertama dan lamanya menabung (dalam minggu).
		F.S. total tabungan kakak setelah n minggu.
	*/

	totalTabungan := tabPertama

	for i := 2; i <= n; i++ {
		totalTabungan += tabPertama + 2500
		tabPertama += 2500
	}

	return totalTabungan
}
