func penjumlahan(N int) {
	var hasil int
	hasil = 0
	for i := 1; i <= N * 10; i++ {
	hasil += i
	}
	fmt.Println(hasil)
   }
   