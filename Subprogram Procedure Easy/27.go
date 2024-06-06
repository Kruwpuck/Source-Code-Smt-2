func identifikasiBaris(u1, u2, u3 int) {
	/*
	I.S. terdefinisi tiga buah bilangan bulat.
	F.S. tercetak string "ganjil" jika termasuk barisan ganjil, string "genap"
	*/
	var hasil string
	if u1%2 == 1 && u2%2 == 1 && u3%2 == 1 {
	hasil = "ganjil"
	} else if u1%2 == 0 && u2%2 == 0 && u3%2 == 0 {
	hasil = "genap"
	} else if u1%2 == 0 || u1%2 == 1 && u2%2 == 0 || u2%2 == 1 && u3%2 == 0 || u3%2 == 0{
	hasil = "bukan ganjil atau genap"
	}
	fmt.Println(hasil)
}
   