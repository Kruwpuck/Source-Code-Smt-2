func findMax(a, b, c, terbesar int) {
	/*
	   IS : Terdefinisi 3 bilangan bulat a, b, dan c
	   FS : Menampilkan bilangan terbesar
	*/
	fmt.Scan(&a, &b, &c)
	terbesar = a
	if b > terbesar { // pengkondisian nilai terbesar antar variabel
		terbesar = b // assign nilai
	}
	if c > terbesar { // pengkondisian nilai terbesar antar variabel
		terbesar = c // assign nilai
	}
	fmt.Println(terbesar)
}
