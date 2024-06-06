func identifikasiBaris(u1, u2, u3 int) {
	/*
	I.S. terdefinisi tiga buah bilangan bulat.
	F.S. tercetak "aritmatika" jika termasuk barisan aritmatika, "geometri" jika termas
	*/
	 var hasil string
	 if u2-u1 == u3-u2 {
	 hasil = "aritmatika"
	 } else if u2/u1 == u3/u2 {
	 hasil = "geometri"
	 } else {
	 hasil = "bukan aritmatika atau geometri"
	 }
	 fmt.Println(hasil)
	}
	