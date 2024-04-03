// Buatlah fungsi saja

func lastThreeChar(last string) string {
	// Mengembalikan 3 karakter terakhir dari string
	var panjangTeks int
	panjangTeks = len(last)
	return last[panjangTeks-3 : panjangTeks]
}