func penjumlahanInterval(M int, N int) {
	/*
	I.S. terdefinisi dua buah bilangan bulat M dan N
	F.S. menampilkan hasil penjumlahan dari M hingga N
	*/
	var hasil int
	hasil = 0
	for i := M; i <= N; i++ {
	   hasil += i
	}
	fmt.Println(hasil)
   }
   