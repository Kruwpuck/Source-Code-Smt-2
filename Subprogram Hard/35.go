func permutasi(n, r int) int {
	/* I.S. Terdefinisi nilai bilangan bulat n dan r sebagai input parameter.
	F.S. Fungsi mengembalikan hasil perhitungan bilangan permutasi dari n dan r */

	if r == 0 {
		return 1
	} else {
		return n * permutasi(n-1, r-1)
	}
}
