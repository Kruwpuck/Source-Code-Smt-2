func barisAritmatika(u1, u2, u3 int) {
	/*
	I.S. terdefinisi tiga buah bilangan bulat.
	F.S. tercetak karakter 'y' jika termasuk barisan aritmatika atau 't' jika t
	*/
	var hasil string
	if u2-u1 == u3-u2 {
	hasil = "ya"
	} else {
	hasil = "tidak"
	}
	fmt.Print(hasil)
   }