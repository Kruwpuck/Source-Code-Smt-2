func deret(n int)float64{
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif n.
	F.S. Fungsi mengembalikan hasil nilai dari deret bilangan }*/
	if n == 1 {
	return 1
	} else {
	return deret(n-1) + 1/float64(n) // gunakan baris ini untuk memproses perh
	}
   }