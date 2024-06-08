func hitungJam(N int) int {
	/*I.S. Terdefinisi nilai bilangan bulat x.
	  /*F.S. fungsu mengembalikan jumlah bernilai integer*/
	if N <= 1 {
		return -1
	} else {
		return 1 + hitungJam(N/2)
	}
}