func telepati(a, b, c int) {
    /* I.S. Terdefinisi a b dan c sebagai bilangan bulan positif 
       F.S. menampilkan dua nilai yang menyatakan banyaknya pemain yang bersahabat dan tidak.
    
    */ 
    
	var sahabat, notsahabat int

	sahabat = 0
	notsahabat = 0
	fmt.Scan(&a, &b, &c)

	for a != 0 && b!= 0 && c != 0{ // lakukan kondisi pengulangan untuk memperoleh jumlah
		if a != b && a != c && ( a + b == c || a + c == b || b + c == a) {
			sahabat++
		} else {
			notsahabat++
		}
		fmt.Scan(&a, &b, &c)
	}
	fmt.Print(sahabat, notsahabat) // tampilkan hasil 2 nilai yang menyatakan banyaknya pemain yang bersahabat dan tidak.
}