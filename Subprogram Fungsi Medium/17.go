// Buat fungsi saja
func segitigaSamaKaki(a, b, c float64) bool {
	// Mengembalikan nilai boolean jika segitiga sama kaki
	var hasil bool
	hasil = (a == b || a == c) && a != c
	return hasil
   }
   