func faktor(i, N int) {
	if i <= N {
	 // pengkondisian untuk mengecek nilai bilangan genap / ganjil
	 if N % i == 0 {
	 fmt.Print(i, " ")
	 }
	 faktor(i+1, N) // tuliskan fungsi rekursif pada baris ini sehingga output
	 }
	}