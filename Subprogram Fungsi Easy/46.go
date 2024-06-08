func hitungJumlah(a, b int) int {
	/* function mengembalikan jumlah dua bilangan, jika kedua bilangan itu ganjil
	   Jika tidak ganjil keduanya, kembalikan bilangan 0
	*/
	var jumlah int
	var ganjil bool
	ganjil = a%2 != 0 && b%2 != 0
	if ganjil {
		jumlah = a + b
	}
	return jumlah
}