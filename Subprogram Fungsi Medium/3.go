func warna(warna1, warna2 string) string {
	/* mengembalikan hasil pencampuran dua warna dasar menjadi warna sekunder */
	var hasil string

	if (warna1 == "merah" && warna2 == "kuning") || (warna1 == "kuning" && warna2 == "merah") {
		hasil = "orange"
	} else if (warna1 == "kuning" && warna2 == "biru") || (warna1 == "biru" && warna2 == "kuning") {
		hasil = "hijau"
	} else if (warna1 == "biru" && warna2 == "merah") || (warna1 == "merah" && warna2 == "biru") {
		hasil = "ungu"
	} else {
		hasil = "Warna tidak valid"
	}

	return hasil
}