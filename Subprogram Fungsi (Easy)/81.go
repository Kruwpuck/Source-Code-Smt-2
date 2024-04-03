// Buatlah fungsi saja

func lastChar(akhir string) byte {
	// Mengembalikan karakter yang berada di indeks terakhir dari string
	var panjangTeks int
	panjangTeks = len(akhir)
	return akhir[panjangTeks-1]
}