// Buatlah fungsi saja

func pemenang(adik, kakak int) string {
	/*
		Fungsi mengembalikan nilai string "adik" jika adik sebagai pemenang
		atau string "kakak" jika kakak sebagai pemenang.
		Adik dinyatakan sebagai pemenang jika tebakannya sama dengan tebakan kakak, atau
		jika selisih tebakan adik dengan kakak 1 atau 5. Jika tidak, maka kakak sebagai
		pemenang.
	*/
	var adikMenang bool
	adikMenang = adik == kakak || kakak == adik || adik-kakak == 1 || adik-kakak == 5 || kakak-adik == 1 || kakak-adik == 5
	if adikMenang {
		return "adik"
	} else {
		return "kakak"
	}

}