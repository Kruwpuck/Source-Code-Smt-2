func hitungPerkalian(a, b int) int {
	/* function mengembalikan hasil perkalian 2 bilangan, jika kedua bilangan genap
	   kembalikan bilangan 0 jika keduanya tidak genap
	*/
	var jumlah int
	var genap bool
	genap = a%2 == 0 && b%2 == 0
	if genap {
		jumlah = a * b
	}
	return jumlah
}
