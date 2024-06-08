func descending(a, b, c int) {
	/* I.S. terdefinisi 3 bilangan bulat a, b, dan c
	   F.S. program mengeluarkan 3 bilangan yang sudah terurut secara mengecil (descending) */

	// Menggunakan bubble sort untuk mengurutkan tiga bilangan secara descending
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}

	// Cetak bilangan yang sudah terurut secara descending
	fmt.Println(a, b, c)
}