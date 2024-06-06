//buat prosedur saja
func berlibur(naikGaji, dapatBonus bool) {
	/*
	I.S. terdefinisi dua buah nilai boolean yang menyatakan status gaji (True bila naik
	F.S. tercetak string "berlibur" jika karyawan naik gaji atau dapat bonus, atau str
	*/
	 var hasil string
	 if naikGaji || dapatBonus == true {
	 hasil = "berlibur"
	 } else {
	 hasil = "tidak berlibur"
	 }
	 fmt.Println(hasil)
	} 