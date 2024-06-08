func kombinasi(n, r int) int {
	/* I.S. Terdefinisi nilai bilangan bulat n dan r sebagai input parameter.
	F.S. Fungsi mengembalikan hasil perhitungan bilangan kombinasi dari n dan r */

	if r == 0 || n == r {
		return 1
	} else {
		return kombinasi(n-1, r) + kombinasi(n-1, r-1)
	}
}