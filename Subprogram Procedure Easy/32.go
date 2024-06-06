//buatlah prosedur saja
func hitungJumlah(n int, hasil *int) {
	/*
	I.S. Terdefinisi sebuah bilangan bulat N
	F.S. S berisi nilai hasil penjumlahan dari 1 hingga N
	*/
	*hasil = 0
	for i := 1; i <= n; i++ {
	*hasil += i
	}
   }